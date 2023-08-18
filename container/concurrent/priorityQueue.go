package concurrent

import (
	"container/heap"
	"github.com/shura1014/common/container/queue"
	"math"
	"sync"
)

// PriorityQueue
// 这是一个堆结构，高的优先级的queue在低优先级之前，每次获取的都是优先级较高的queue
type PriorityQueue struct {
	mu           sync.Mutex
	heap         *queue.PriorityQueue
	nextPriority int64
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		heap:         queue.NewPriorityQueue(),
		nextPriority: math.MaxInt64,
	}
	// 要是一个堆，需要实现堆接口
	//type Interface interface {
	//	sort.Interface
	//	Push(x any) // add x as element Len()
	//	Pop() any   // remove and return element Len() - 1.
	//}
	heap.Init(pq.heap)
	return pq
}

func (pq *PriorityQueue) NextPriority() int64 {
	return pq.nextPriority
}

func (pq *PriorityQueue) Push(value any, priority int64) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	heap.Push(pq.heap, &queue.PriorityQueueItem{
		Value: value, Priority: priority,
	})

	nextPriority := pq.nextPriority
	if priority >= nextPriority {
		return
	}
	pq.nextPriority = priority
}

func (pq *PriorityQueue) Pop() any {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if v := heap.Pop(pq.heap); v != nil {
		var nextPriority int64 = math.MaxInt64
		if len(*pq.heap) > 0 {
			nextPriority = (*pq.heap)[0].Priority
		}
		pq.nextPriority = nextPriority
		return v.(*queue.PriorityQueueItem).Value
	}
	return nil
}
