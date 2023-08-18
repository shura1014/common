package concurrent

import (
	"github.com/shura1014/common/container/stack"
	"sync"
)

// Stack 栈 先进后出
type Stack[V any] struct {
	mu    sync.RWMutex
	stack *stack.Stack[V]
}

func NewStack[V any]() *Stack[V] {
	return &Stack[V]{stack: stack.NewStack[V]()}
}

func (s *Stack[V]) Push(e V) {
	s.mu.Lock()
	s.stack.Push(e)
	s.mu.Unlock()
}

// Pop 从 s 中取出最后放入 栈 的值
func (s *Stack[V]) Pop() V {
	s.mu.Lock()
	pop := s.stack.Pop()
	s.mu.Unlock()
	return pop
}

func (s *Stack[V]) Peek() V {
	s.mu.RLock()
	peek := s.stack.Peek()
	s.mu.RUnlock()
	return peek
}

// Len 返回 栈 的长度
func (s *Stack[V]) Len() int {
	s.mu.RLock()
	i := s.stack.Len()
	s.mu.RUnlock()
	return i
}

// IsEmpty 判空
func (s *Stack[V]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.stack.IsEmpty()
}

func (s *Stack[V]) Iterator(f func(v any)) {
	s.mu.RLock()
	s.stack.Iterator(func(v any) {
		f(v)
	})
	s.mu.RUnlock()
}
