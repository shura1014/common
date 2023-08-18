package stack

// Stack 栈 先进后出
type Stack[V any] struct {
	element []V
}

func NewStack[V any]() *Stack[V] {
	return &Stack[V]{element: []V{}}
}

func (s *Stack[V]) Push(e V) {
	s.element = append(s.element, e)
}

// Pop 从 s 中取出最后放入 栈 的值
func (s *Stack[V]) Pop() V {
	res := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return res
}

// Peek 获取但不删除
func (s *Stack[V]) Peek() V {
	res := s.element[len(s.element)-1]
	return res
}

// Len 返回 栈 的长度
func (s *Stack[V]) Len() int {
	return len(s.element)
}

// IsEmpty 判空
func (s *Stack[V]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[V]) Iterator(f func(v any)) {
	for i := len(s.element) - 1; i >= 0; i-- {
		f(s.Pop())
	}
}
