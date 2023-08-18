package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("1")
	stack.Push("3")
	stack.Push("2")
	stack.Push("8")
	stack.Push("5")
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())

}

func TestStack_Iterator(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("1")
	stack.Push("3")
	stack.Push("2")
	stack.Push("8")
	stack.Push("5")

	stack.Iterator(func(v any) {
		t.Log(v)
	})
}
