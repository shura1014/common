package random

import (
	"crypto/rand"
	"github.com/shura1014/common/assert"
	"io"
	"testing"
)

// BenchmarkByte-8   	 4004636	       274.5 ns/op
func BenchmarkByte(b *testing.B) {

	for i := 0; i < b.N; i++ {
		k := make([]byte, 4)
		if _, err := io.ReadFull(rand.Reader, k); err != nil {

		}
	}
}

// BenchmarkCacheByte-8   	12702120	        87.81 ns/op
func BenchmarkCacheByte(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Byte(4)
	}
}

func BenchmarkCacheInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		n := Int(10000)
		assert.Assert(n >= 0 && n <= 10000, true)
	}
}

func BenchmarkCacheIntRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := IntRange(-100, 100)
		assert.Assert(n >= -100 && n <= 100, true)
	}

	for i := 0; i < b.N; i++ {
		n := IntRange(100, 10000)
		assert.Assert(n >= 100 && n <= 10000, true)
	}
}

func TestCacheByte(t *testing.T) {
	bytes := Byte(5)
	t.Log(bytes)
}

func TestCacheInt(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := Int(100)
		t.Log(n)
	}
}

func TestCacheIntRange(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := IntRange(-100, 100)
		t.Log(n)
	}
}

func TestCacheSymbols(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := Symbols(6)
		t.Log(n)
	}
}

func TestCacheLetter(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := Letter(6)
		t.Log(n)
	}
}

func TestCacheW(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := W(10)
		t.Log(n)
	}
}

func TestCacheChar(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := Char(16)
		t.Log(n)
	}

	//t.Log(len(characters))
}

func TestCacheNumber(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := Number(6)
		t.Log(n)
	}
}
