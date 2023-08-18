package byteutil

import (
	"testing"
)

func TestToByte(t *testing.T) {
	type T struct {
		data any
	}

	t.Log(ToByte(&T{}))
	t.Log(ToByte("11111"))

}
