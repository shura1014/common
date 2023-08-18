package queue

import (
	"errors"
	"github.com/shura1014/common/type/atom"
)

// BlockingQueue 阻塞队列
type BlockingQueue struct {
	limit  int
	closed *atom.Bool
	// 通道本身就是一个先进先出的队列
	queue chan any
}

func NewBlockingQueue(limit int) *BlockingQueue {
	return &BlockingQueue{
		limit:  limit,
		closed: atom.NewBool(),
		queue:  make(chan any, limit),
	}
}

func (q *BlockingQueue) IsClosed() bool {
	return q.closed.Load()
}

func (q *BlockingQueue) Push(e any) error {
	if q.IsClosed() {
		return errors.New("chan is closed")
	}
	q.queue <- e
	return nil
}

func (q *BlockingQueue) Pop() (e any, err error) {
	if q.IsClosed() {
		return nil, errors.New("chan is closed")
	}
	e = <-q.queue
	return e, nil
}

func (q *BlockingQueue) Close() {
	if !q.closed.Cas(false, true) {
		return
	}
	defer func() {
		if q.closed.Load() {
			_ = recover()
		}
	}()
	close(q.queue)
}

func (q *BlockingQueue) Size() (length int64) {
	return int64(len(q.queue))
}
