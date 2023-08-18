package concurrent

import (
	"github.com/shura1014/common/container/set"
	"sync"
)

type Set struct {
	mu  sync.RWMutex
	set *set.Set
}

func NewSet() *Set {
	return &Set{
		set: set.NewSet(),
	}
}

func (s *Set) Add(items ...any) {
	s.mu.Lock()
	s.set.Add(items...)
	s.mu.Unlock()
}

func (s *Set) Remove(item any) {
	s.mu.Lock()
	s.set.Remove(item)
	s.mu.Unlock()
}

func (s *Set) Clear() {
	s.mu.Lock()
	s.set.Clear()
	s.mu.Unlock()
}

func (s *Set) Size() int {
	s.mu.RLock()
	size := s.set.Size()
	s.mu.RUnlock()
	return size
}

func (s *Set) Contains(item any) bool {
	s.mu.RLock()
	ok := s.set.Contains(item)
	s.mu.RUnlock()
	return ok
}

func (s *Set) Iterator(f func(v interface{})) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Iterator(func(v interface{}) {
		f(v)
	})
}
