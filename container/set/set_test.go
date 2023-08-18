package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add("1")
	set.Add("1")
	set.Add("1")
	set.Add("2")
	set.Add("1")
	set.Add("3")
	set.Add("1")
	set.Add("1")
	set.Add("1")
	set.Add("5")

	set.Iterator(func(v interface{}) {
		t.Log(v)
	})
}
