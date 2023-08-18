package queue

// Queue 是用于存放 int 的队列
type Queue[V any] struct {
	element []V
}

// NewQueue 返回 *kit.Queue
func NewQueue[V any]() *Queue[V] {
	return &Queue[V]{element: []V{}}
}

// Push 把 element 放入队列
func (q *Queue[V]) Push(e V) {
	q.element = append(q.element, e)
}

// Pop 从 element 中取出最先进入队列的值
func (q *Queue[V]) Pop() V {
	res := q.element[0]
	q.element = q.element[1:]
	return res
}

// Peek 获取但不删除
func (q *Queue[V]) Peek() V {
	res := q.element[0]
	return res
}

// Len 返回 element 的长度
func (q *Queue[V]) Len() int {
	return len(q.element)
}

// IsEmpty 判空
func (q *Queue[V]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue[V]) Iterator(f func(v any)) {
	for i := 0; i < len(q.element); i++ {
		f(q.element[i])
	}
}
