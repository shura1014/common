package concurrent

import (
	"sync"
)

type Map[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V)}
}
func (m *Map[K, V]) Get(key K) (value V) {
	m.mu.RLock()
	value, _ = m.data[key]
	m.mu.RUnlock()
	return

}

func (m *Map[K, V]) Contains(key K) (ok bool) {
	m.mu.RLock()
	_, ok = m.data[key]
	m.mu.RUnlock()
	return

}

func (m *Map[K, V]) Put(key K, value V) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()

}

func (m *Map[K, V]) PutAll(data map[K]V) {
	m.mu.Lock()
	for key, value := range data {
		m.data[key] = value
	}
	m.mu.Unlock()
}

func (m *Map[K, V]) Remove(keys ...K) {
	m.mu.Lock()
	for _, k := range keys {
		delete(m.data, k)
	}
	m.mu.Unlock()
}

func (m *Map[K, V]) Iterator(f func(key K, value V) bool) {
	m.mu.RLock()
	for k, v := range m.data {
		if f(k, v) {
			break
		}
	}
	m.mu.RUnlock()
}

func (m *Map[K, V]) GetAll() map[K]V {
	m.mu.RLock()
	data := make(map[K]V)
	m.Iterator(func(key K, value V) (ok bool) {
		data[key] = value
		return
	})
	m.mu.RUnlock()
	return data

}
