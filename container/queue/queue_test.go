package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[string]()
	queue.Push("1")
	queue.Push("3")
	queue.Push("2")
	queue.Push("8")
	queue.Push("5")

	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Pop()
}

func TestQueue_Iterator(t *testing.T) {
	queue := NewQueue[string]()
	queue.Push("1")
	queue.Push("3")
	queue.Push("2")
	queue.Push("8")
	queue.Push("5")
	queue.Iterator(func(v any) {
		t.Log(v)
	})
}
