package hash

import "testing"

func TestBKDRHash64(t *testing.T) {
	hash64 := BKDRHash64([]byte("123"))
	t.Log(hash64)
}
