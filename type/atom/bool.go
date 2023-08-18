package atom

import "sync/atomic"

type Bool struct {
	value int32
}

func NewBool(value ...bool) *Bool {
	t := &Bool{}
	if len(value) > 0 {
		if value[0] {
			t.value = 1
		} else {
			t.value = 0
		}
	}
	return t
}

func (v *Bool) Swap(value bool) (old bool) {
	if value {
		old = atomic.SwapInt32(&v.value, 1) == 1
	} else {
		old = atomic.SwapInt32(&v.value, 0) == 1
	}
	return
}

func (v *Bool) Cas(old, new bool) (swapped bool) {
	var oldInt32, newInt32 int32
	if old {
		oldInt32 = 1
	}
	if new {
		newInt32 = 1
	}
	return atomic.CompareAndSwapInt32(&v.value, oldInt32, newInt32)
}

func (v *Bool) Load() bool {
	return atomic.LoadInt32(&v.value) > 0
}

func (v *Bool) String() string {
	if v.Load() {
		return "true"
	}
	return "false"
}
