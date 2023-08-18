package queue

type PriorityQueueItem struct {
	Value    any
	Priority int64
}

// PriorityQueue 是一个堆结构
// 一个堆结构需要实现 Push Pop Len Less Swap 方法
type PriorityQueue []*PriorityQueueItem

func NewPriorityQueue() *PriorityQueue {
	queue := make(PriorityQueue, 0)
	return &queue
}
func (p *PriorityQueue) Len() int {
	return len(*p)
}

// Less 最小的被放在堆的顶部
func (p *PriorityQueue) Less(i, j int) bool {
	return (*p)[i].Priority < (*p)[j].Priority
}

func (p *PriorityQueue) Swap(i, j int) {
	if len(*p) == 0 {
		return
	}
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

// Push 放置最后
func (p *PriorityQueue) Push(queueItem any) {
	*p = append(*p, queueItem.(*PriorityQueueItem))
}

// Pop 取数组的最后一个
func (p *PriorityQueue) Pop() any {
	length := len(*p)
	if length == 0 {
		return nil
	}
	item := (*p)[length-1]
	*p = (*p)[0 : length-1]
	return item
}
