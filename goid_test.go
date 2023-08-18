package common

import (
	"testing"
)

func TestId(t *testing.T) {
	for i := 0; i < 1; i++ {
		go func() {
			t.Log(GoID())
			t.Log(GoID())
			t.Log(GoID())
			t.Log(GoID())
			t.Log(GoID())
			t.Log(GoID())
			t.Log(GoID())
			t.Log(GoID())
		}()
	}
	Wait()
}
