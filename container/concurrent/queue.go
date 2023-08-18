package concurrent

import (
	"github.com/shura1014/common/container/queue"
	"sync"
)

type Queue[V any] struct {
	mu    sync.RWMutex
	queue *queue.Queue[V]
}

func NewQueue[V any]() *Queue[V] {
	return &Queue[V]{queue: queue.NewQueue[V]()}
}

func (q *Queue[V]) Push(e V) {
	q.mu.Lock()
	q.queue.Push(e)
	q.mu.Unlock()
}

func (q *Queue[V]) Peek() V {
	q.mu.RLock()
	peek := q.queue.Peek()
	q.mu.RUnlock()
	return peek
}

// Pop 从 element 中取出最先进入队列的值
func (q *Queue[V]) Pop() V {
	q.mu.Lock()
	pop := q.queue.Pop()
	q.mu.Unlock()
	return pop
}

// Len 返回 element 的长度
func (q *Queue[V]) Len() int {
	q.mu.RLock()
	s := q.queue.Len()
	q.mu.RUnlock()
	return s
}

func (q *Queue[V]) IsEmpty() bool {
	return q.queue.IsEmpty()
}

func (q *Queue[V]) Iterator(f func(v any)) {
	q.mu.RLock()
	q.queue.Iterator(func(v any) {
		f(v)
	})
	q.mu.RUnlock()
}
