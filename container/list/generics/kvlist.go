package generics

// MapElement Use list.Element 我们给它加上一个泛型
// 专门为lru缓存定制的list，能拿到K Y,这样使得实现lru的时候不用额外定义有一个entry{k,v}
type MapElement[K comparable, V any] struct {
	next, prev *MapElement[K, V]

	list *MapList[K, V]

	Key   K
	Value V
}

func (e *MapElement[K, V]) Next() *MapElement[K, V] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *MapElement[K, V]) Prev() *MapElement[K, V] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type MapList[K comparable, V any] struct {
	root MapElement[K, V]
	len  int
}

func (l *MapList[K, V]) Init() *MapList[K, V] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func NewMapList[K comparable, V any]() *MapList[K, V] { return new(MapList[K, V]).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *MapList[K, V]) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *MapList[K, V]) Front() *MapElement[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *MapList[K, V]) Back() *MapElement[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *MapList[K, V]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *MapList[K, V]) insert(e, at *MapElement[K, V]) *MapElement[K, V] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *MapList[K, V]) insertValue(k K, v V, at *MapElement[K, V]) *MapElement[K, V] {
	return l.insert(&MapElement[K, V]{Key: k, Value: v}, at)
}

// remove removes e from its list, decrements l.len
func (l *MapList[K, V]) remove(e *MapElement[K, V]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}

// move moves e to next to at.
func (l *MapList[K, V]) move(e, at *MapElement[K, V]) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *MapList[K, V]) Remove(e *MapElement[K, V]) any {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *MapList[K, V]) PushFront(k K, v V) *MapElement[K, V] {
	l.lazyInit()
	return l.insertValue(k, v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *MapList[K, V]) PushBack(k K, v V) *MapElement[K, V] {
	l.lazyInit()
	return l.insertValue(k, v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *MapList[K, V]) InsertBefore(k K, v V, mark *MapElement[K, V]) *MapElement[K, V] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(k, v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *MapList[K, V]) InsertAfter(k K, v V, mark *MapElement[K, V]) *MapElement[K, V] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(k, v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *MapList[K, V]) MoveToFront(e *MapElement[K, V]) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *MapList[K, V]) MoveToBack(e *MapElement[K, V]) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *MapList[K, V]) MoveBefore(e, mark *MapElement[K, V]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *MapList[K, V]) MoveAfter(e, mark *MapElement[K, V]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *MapList[K, V]) PushBackList(other *MapList[K, V]) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Key, e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *MapList[K, V]) PushFrontList(other *MapList[K, V]) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Key, e.Value, &l.root)
	}
}
