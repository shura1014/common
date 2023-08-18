package atom

import (
	"strconv"
	"sync/atomic"
)

type Int64 struct {
	value int64
}

func NewInt64(value ...int64) *Int64 {
	if len(value) > 0 {
		return &Int64{
			value: value[0],
		}
	}
	return &Int64{}
}

func (v *Int64) Swap(value int64) (old int64) {
	return atomic.SwapInt64(&v.value, value)
}

func (v *Int64) Load() int64 {
	return atomic.LoadInt64(&v.value)
}

func (v *Int64) Add(delta int64) (new int64) {
	return atomic.AddInt64(&v.value, delta)
}

func (v *Int64) Cas(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&v.value, old, new)
}

func (v *Int64) String() string {
	return strconv.FormatInt(v.Load(), 10)
}
