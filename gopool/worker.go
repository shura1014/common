package gopool

import (
	"github.com/shura1014/common/goerr"
	"time"
)

type Worker struct {
	pool *GoPool
	// 任务队列
	task chan *Task
	// 最后执行任务的时间
	lastTime time.Time
}

func (w *Worker) run() {
	w.pool.incrRunning()
	go w.running()
}

// running
// 真正执行业务逻辑的地方
// f()
// 用完之后记得归还worker，让其它任务可以执行
// 执行过程中如果发生异常应该捕获（执行异常处理逻辑，用户可以自定义），并且归还worker
func (w *Worker) running() {
	var useTask Task
	// 发生异常
	defer func() {
		w.pool.decRunning()
		w.pool.workerPool.Put(w)
		if err := recover(); err != nil {
			w.pool.PanicHandle(useTask.ctx, goerr.Wrap(err))
		}
		w.pool.cond.Signal()
	}()
	for task := range w.task {
		if task == nil {
			w.pool.workerPool.Put(w)
			return
		}
		useTask = *task
		task.taskFunc(task.ctx)
		// 任务运行完成 worker空闲还回去
		w.pool.recycleWorker(w)
		//w.pool.decRunning()
	}
}
