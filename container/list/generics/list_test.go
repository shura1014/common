package generics

import "testing"

func TestVList(t *testing.T) {
	type User struct {
		name string
		age  int
	}
	list := NewList[User]()
	list.PushFront(User{name: "name1", age: 10})
	list.PushFront(User{name: "name2", age: 20})
	list.PushFront(User{name: "name3", age: 20})

	t.Logf(list.Front().Value.name)
	t.Logf(list.Back().Value.name)
}

func TestKVList(t *testing.T) {
	type User struct {
		name string
		age  int
	}
	list := NewMapList[string, User]()
	list.PushFront("key1", User{name: "name1", age: 10})
	list.PushFront("key2", User{name: "name2", age: 20})
	list.PushFront("key3", User{name: "name3", age: 20})
	t.Logf(list.Front().Key)
	t.Logf(list.Front().Value.name)
}
