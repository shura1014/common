package objectutil

import (
	"testing"
	"time"
)

func TestIsZero(t *testing.T) {
	t.Log(IsZero(0))
	t.Log(IsZero(1))
	t.Log(IsZero(""))
	t.Log(IsZero("1"))
	t.Log(IsZero(time.Duration(0)))
	t.Log(IsZero(time.Duration(1)))
	t.Log(IsZero(nil))
	t.Log(IsZero(make(map[string]any)))
}
