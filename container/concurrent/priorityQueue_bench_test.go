package concurrent

import (
	"testing"
)

// Benchmark_PriorityQueue-8   	 8608009	       142.4 ns/op
// Benchmark_PriorityQueue-8   	 8332312	       134.3 ns/op
func Benchmark_PriorityQueue_Push(b *testing.B) {
	queue := NewPriorityQueue()
	var i int64
	b.RunParallel(func(pb *testing.PB) {
		i = 0
		for pb.Next() {
			queue.Push("test10", i)
			i++
		}
	})
}

// Benchmark_PriorityQueue_POP-8   	31500954	        41.28 ns/op
func Benchmark_PriorityQueue_POP(b *testing.B) {
	queue := NewPriorityQueue()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			queue.Pop()
		}
	})
}
