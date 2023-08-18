package concurrent

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pQueue := NewPriorityQueue()

	pQueue.Push("test10", 10)
	pQueue.Push("test2", 2)
	pQueue.Push("test8", 8)
	pQueue.Push("test11", 11)
	pQueue.Push("test5", 5)

	t.Log(pQueue.nextPriority)
	t.Log(pQueue.Pop())
	t.Log(pQueue.nextPriority)

	t.Log(pQueue.Pop())
	t.Log(pQueue.nextPriority)

	t.Log(pQueue.Pop())
	t.Log(pQueue.nextPriority)

	t.Log(pQueue.Pop())
	t.Log(pQueue.nextPriority)

	t.Log(pQueue.Pop())

}
