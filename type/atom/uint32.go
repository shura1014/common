package atom

import (
	"strconv"
	"sync/atomic"
)

type Uint32 struct {
	value uint32
}

func NewUint32(value ...uint32) *Uint32 {
	if len(value) > 0 {
		return &Uint32{
			value: value[0],
		}
	}
	return &Uint32{}
}

func (v *Uint32) Swap(value uint32) (old uint32) {
	return atomic.SwapUint32(&v.value, value)
}

func (v *Uint32) Load() uint32 {
	return atomic.LoadUint32(&v.value)
}

func (v *Uint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&v.value, delta)
}

func (v *Uint32) Cas(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&v.value, old, new)
}

func (v *Uint32) String() string {
	return strconv.Itoa(int(v.Load()))
}
