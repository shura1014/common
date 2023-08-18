package queue

import (
	"testing"
	"time"
)

func TestBlockingQueue(t *testing.T) {
	queue := NewBlockingQueue(100)
	go func() {
		for i := 0; i < 10000; i++ {
			_ = queue.Push(i)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			for j := i * 1000; j < i*1000+1000; i++ {
				pop, _ := queue.Pop()
				if i != pop {
					t.Fail()
					return
				}
			}
		}

	}()
	time.Sleep(time.Second * 8)
	//utils.Wait()
}
