package concurrent

import (
	"github.com/shura1014/common/container/list/generics"
	"github.com/shura1014/common/utils/stringutil"
	"sync"
)

type AnyMap[V any] struct {
	mu   sync.RWMutex
	data map[any]V
}

func NewAnyMap[V any]() *AnyMap[V] {
	return &AnyMap[V]{data: make(map[any]V)}
}
func (m *AnyMap[V]) Get(key any) (value V) {
	m.mu.RLock()
	value, _ = m.data[key]
	m.mu.RUnlock()
	return

}

func (m *AnyMap[V]) Contains(key any) (ok bool) {
	m.mu.RLock()
	_, ok = m.data[key]
	m.mu.RUnlock()
	return

}

func (m *AnyMap[V]) Put(key any, value V) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()

}

func (m *AnyMap[V]) PutAll(data map[any]V) {
	m.mu.Lock()
	for key, value := range data {
		m.data[key] = value
	}
	m.mu.Unlock()
}

func (m *AnyMap[V]) Remove(keys ...any) {
	m.mu.Lock()
	for _, k := range keys {
		delete(m.data, k)
	}
	m.mu.Unlock()
}

func (m *AnyMap[V]) Iterator(f func(key any, value V) bool) {
	m.mu.RLock()
	for k, v := range m.data {
		if f(k, v) {
			break
		}
	}
	m.mu.RUnlock()
}

func (m *AnyMap[V]) KeysArray() []any {
	var keys []any
	m.Iterator(func(key any, value V) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

func (m *AnyMap[V]) KeysList() *generics.List[any] {
	l := generics.NewList[any]()
	m.Iterator(func(key any, value V) bool {
		l.PushBack(key)
		return true
	})
	return l
}

func (m *AnyMap[V]) ValuesArray() []V {
	var values []V
	m.Iterator(func(key any, value V) bool {
		values = append(values, value)
		return true
	})
	return values
}

func (m *AnyMap[V]) ValuesList() *generics.List[V] {
	l := generics.NewList[V]()
	m.Iterator(func(key any, value V) bool {
		l.PushBack(value)
		return true
	})
	return l
}

func (m *AnyMap[V]) GetAll() map[any]V {
	m.mu.RLock()
	data := make(map[any]V)
	m.Iterator(func(key any, value V) (ok bool) {
		data[key] = value
		return
	})
	m.mu.RUnlock()
	return data

}

func (m *AnyMap[V]) String() string {
	return stringutil.ToString(m.data)
}
