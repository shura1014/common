package binaryutil

import (
	"testing"
)

func TestToByte(t *testing.T) {
	data := []any{"1", 2, false, []int32{1, 2, 3}, map[string]any{"name": "wendell"}}
	toByte := Encode(data)
	t.Log(toByte)
	decode := make([]byte, len(toByte))
	//decode := DecodeToString(toByte)
	_ = Decode(toByte, decode)
	t.Log(string(decode))
}
