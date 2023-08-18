package atom

import "testing"

func BenchmarkAtom(b *testing.B) {
	num := NewInt64()
	for i := 0; i < b.N; i++ {
		num.Add(1)
	}

	b.Log(num.Load())
}

func TestAtom(t *testing.T) {
	num := NewInt64()
	load := num.Load()
	num.Cas(load, load+1)

	num.Swap(100)
}
