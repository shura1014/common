package stringutil

import (
	"testing"
	"time"
)

// 驼峰转下划线
func TestUnderline(t *testing.T) {
	t.Log(Underline("AbstractBalance"))
	t.Log(LowUnderline("AbstractBalance"))
	t.Log(UpperUnderline("AbstractBalance"))
}

func TestJoin(t *testing.T) {
	t.Log(Join("a", "b", "c")) // anc
}

func TestIsArray(t *testing.T) {
	t.Log(IsArray([]string{"a", "b"}, "a")) // true
}

func TestToString(t *testing.T) {
	t.Log(ToString([]byte{'1', '2'})) // 12
	t.Log(ToString(100))              // 100
	t.Log(ToString(true))             // true
	t.Log(ToString(time.Now()))       // 2023-08-18 14:45:03.316659 +0800 CST m=+0.000671251
}

func TestIsNumeric(t *testing.T) {
	t.Log(IsNumeric("12")) // true
	t.Log(IsNumeric("xx")) // false
}

func TestStringToBytes(t *testing.T) {
	t.Log(StringToBytes("123")) // [49 50 51]
}
