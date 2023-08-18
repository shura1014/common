package atom

import (
	"strconv"
	"sync/atomic"
)

type Int32 struct {
	value int32
}

func NewInt32(value ...int32) *Int32 {
	if len(value) > 0 {
		return &Int32{
			value: value[0],
		}
	}
	return &Int32{}
}

func (v *Int32) Swap(value int32) (old int32) {
	return atomic.SwapInt32(&v.value, value)
}

func (v *Int32) Load() int32 {
	return atomic.LoadInt32(&v.value)
}

func (v *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&v.value, delta)
}

func (v *Int32) Cas(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.value, old, new)
}

func (v *Int32) String() string {
	return strconv.Itoa(int(v.Load()))
}
