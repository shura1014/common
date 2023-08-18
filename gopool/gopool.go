package gopool

import (
	"context"
	"errors"
	"github.com/shura1014/common/clog"
	"github.com/shura1014/common/goerr"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrorInValidCap    = errors.New("cap cannot le zero")
	ErrorInValidExpire = errors.New("expire cannot le zero")
	ErrorPoolIsClosed  = errors.New("pool is closed ")
)

const DefaultExpire = 3

var DefaultHandler = func(ctx context.Context, err *goerr.BizError) {
	clog.Error(err.String())
}

// GoPool 协程池
type GoPool struct {
	// 容量 其实也代表着 workers 的数量
	cap int32

	// 正在运行的数量
	running int32

	// 空闲worker 可以使用的
	workers []*Worker

	//空闲worker超过这个时间回收掉 expireWorker
	expireTime time.Duration

	//// 释放资源信号
	//// 收到此信号，协程池关闭
	//release chan sig

	//lock *lock.Lock
	lock *sync.RWMutex

	// 释放只能调用一次
	shutdown bool

	// 资源池，负责管理worker，创建worker
	workerPool sync.Pool

	// 等待唤醒机制
	cond *sync.Cond

	// 业务执行报错的处理函数
	PanicHandle func(ctx context.Context, err *goerr.BizError)
}

func (p *GoPool) incrRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *GoPool) decRunning() {
	atomic.AddInt32(&p.running, -1)
}

func NewPool(cap int32) (*GoPool, error) {
	return NewTimePool(cap, DefaultExpire)
}

// NewTimePool
// cap 容量
// expire 过期时间 （秒）
func NewTimePool(cap int32, expire int) (*GoPool, error) {
	if cap <= 0 {
		return nil, ErrorInValidCap
	}

	if expire <= 0 {
		return nil, ErrorInValidExpire
	}

	p := &GoPool{
		cap:         cap,
		expireTime:  time.Duration(expire) * time.Second,
		PanicHandle: DefaultHandler,
		//lock:        lock.New(),
		lock: &sync.RWMutex{},
	}

	p.workerPool.New = func() any {
		return &Worker{
			pool: p,
			task: make(chan *Task, 1),
		}
	}
	p.cond = sync.NewCond(p.lock)
	go p.expireWorker()
	return p, nil
}

// addWorker
// 归还worker
// 唤醒阻塞 waitIdleWorker
// 追加到末尾
func (p *GoPool) recycleWorker(w *Worker) {
	w.lastTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, w)
	p.cond.Signal()
	p.lock.Unlock()
}

// Execute
// 新的任务进来，去执行
// 获取执行该任务的Worker
// 将任务给到该Worker
func (p *GoPool) Execute(ctx context.Context, task Func) error {
	if p.IsClosed() {
		return ErrorPoolIsClosed
	}

	w := p.GetWorker()
	t := &Task{
		ctx:      ctx,
		taskFunc: task,
	}
	w.task <- t
	return nil
}

// GetWorker
// 如果有空闲worker的直接取第一个worker返回（因为归还是归还了最后一个）
// 如果没有了需要新建一个worker（需要判断一下是否超出容量）
// 如果超出容量，那么就等待空闲的worker
func (p *GoPool) GetWorker() *Worker {

	// 获取pool的worker
	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers) - 1
	if n >= 0 {
		worker := idleWorkers[0]
		idleWorkers[0] = nil
		p.workers = idleWorkers[1:]
		p.lock.Unlock()
		return worker

	}
	// 如果没有 需要新建一个worker
	if p.running < p.cap {
		p.lock.Unlock()
		c := p.workerPool.Get()
		var w *Worker
		if c == nil {
			w = &Worker{
				pool: p,
				task: make(chan *Task, 1),
			}
		} else {
			w = c.(*Worker)
		}

		w.run()
		return w
	}
	// 如果正在运行的worker大于等于pool容量 阻塞等待worker释放
	p.lock.Unlock()

	return p.waitIdleWorker()
}

// waitIdleWorker
// p.cond.Wait() 等待在此处
// recycleWorker 归还
func (p *GoPool) waitIdleWorker() *Worker {
	p.lock.Lock()
	p.cond.Wait()
	//fmt.Println("得到通知")
	idleWorkers := p.workers
	n := len(idleWorkers) - 1
	if n < 0 {

		// 如果没有 需要新建一个worker
		if p.running < p.cap {
			c := p.workerPool.Get()
			var w *Worker
			if c == nil {
				w = &Worker{
					pool: p,
					task: make(chan *Task, 1),
				}
			} else {
				w = c.(*Worker)
			}

			w.run()
			p.lock.Unlock()
			return w
		}
		p.lock.Unlock()
		// 由于多线程的原因，这里kennel还是拿不到，继续等待
		return p.waitIdleWorker()
	}
	worker := idleWorkers[n]
	idleWorkers[n] = nil
	p.workers = idleWorkers[:n]
	p.lock.Unlock()
	return worker
}

func (p *GoPool) Shutdown() {
	if p.shutdown {
		return
	}
	p.lock.Lock()
	workers := p.workers
	for i, w := range workers {
		w.task = nil
		w.pool = nil
		workers[i] = nil
	}
	p.workers = nil
	p.shutdown = true
	p.lock.Unlock()
}

// expireWorker
// 业务比较空闲，需要处理一些长时间不用的worker
// 由于归还worker是添加到末尾，所以末尾肯定是离过期时间最远的
// 所以从前往后找，只要有一个没有过期，那么后面的就都没有过期，清理前面的即可
func (p *GoPool) expireWorker() {
	ticker := time.NewTicker(p.expireTime)
	for range ticker.C {
		if p.IsClosed() {
			break
		}
		p.lock.Lock()
		idleWorkers := p.workers
		n := len(idleWorkers) - 1
		if n >= 0 {
			var cleanN = -1
			for i, worker := range idleWorkers {
				// 第一个为空闲最久的，如果不满足可以直接return
				if time.Now().Sub(worker.lastTime) <= p.expireTime {
					break
				}
				cleanN = i
				worker.task <- nil
				//idleWorkers[i] = nil
			}
			if cleanN != -1 {
				if cleanN >= len(idleWorkers)-1 {
					p.workers = idleWorkers[:0]
				} else {
					p.workers = idleWorkers[cleanN+1:]
				}
				clog.Debug("清除完成 本次清理数 %d,running:%d\n", cleanN+1, p.running)
			}
		}
		p.lock.Unlock()
	}
}

// Running 当前正在运行的数量
func (p *GoPool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

// Free 可使用数量
func (p *GoPool) Free() int {
	return int(p.cap - p.running)
}

// IsClosed 判断协程池是否已经释放 关闭 close
func (p *GoPool) IsClosed() bool {
	return p.shutdown
}

// Restart
// 如果如果是正在运行状态就不需要重启了
// 如果已经关闭了，那么取消掉关闭信号
func (p *GoPool) Restart() bool {
	if !p.IsClosed() {
		return true
	}
	p.shutdown = false
	go p.expireWorker()
	return true
}
