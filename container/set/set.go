package set

type Set map[any]struct{}

func NewSet() *Set {
	set := make(Set)
	return &set
}

func (s *Set) Add(items ...any) {
	for _, v := range items {
		(*s)[v] = struct{}{}
	}
}

func (s *Set) Remove(item any) {
	if s != nil {
		delete(*s, item)
	}
}

func (s *Set) Clear() {
	*s = make(map[any]struct{})
}

func (s *Set) Size() int {
	return len(*s)
}

func (s *Set) Contains(item any) bool {
	_, ok := (*s)[item]
	return ok
}

func (s *Set) Iterator(f func(v any)) {
	for k := range *s {
		f(k)
	}
}
