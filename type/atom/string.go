package atom

import "sync/atomic"

type String struct {
	value atomic.Value
}

func NewString(value ...string) *String {
	t := &String{}
	if len(value) > 0 {
		t.value.Store(value[0])
	}
	return t
}

func (v *String) Load() string {
	s := v.value.Load()
	if s != nil {
		return s.(string)
	}
	return ""
}

func (v *String) Swap(value string) (old string) {
	old = v.Load()
	v.value.Store(value)
	return
}

func (v *String) String() string {
	return v.Load()
}
