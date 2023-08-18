package atom

import (
	"strconv"
	"sync/atomic"
)

type Uint64 struct {
	value uint64
}

func NewUint64(value ...uint64) *Uint64 {
	if len(value) > 0 {
		return &Uint64{
			value: value[0],
		}
	}
	return &Uint64{}
}

func (v *Uint64) Swap(value uint64) (old uint64) {
	return atomic.SwapUint64(&v.value, value)
}

func (v *Uint64) Load() uint64 {
	return atomic.LoadUint64(&v.value)
}

func (v *Uint64) Add(delta uint64) (new uint64) {
	return atomic.AddUint64(&v.value, delta)
}

func (v *Uint64) Cas(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, old, new)
}

func (v *Uint64) String() string {
	return strconv.Itoa(int(v.Load()))
}
