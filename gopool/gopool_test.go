package gopool

import (
	"context"
	"github.com/shura1014/common"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	// GiB // 1073741824
	// TiB // 1099511627776             (超过了int32的范围)
	// PiB // 1125899906842624
	// EiB // 1152921504606846976
	// ZiB // 1180591620717411303424    (超过了int64的范围)
	// YiB // 1208925819614629174706176
)

const (
	Param    = 100
	PoolSize = 10000
	TestSize = 1000000
	n        = 50000000
)

var curMem uint64

const (
	RunTimes           = 1000000
	BenchParam         = 10
	DefaultExpiredTime = 10 * time.Second
)

func demoFunc() {
	time.Sleep(time.Duration(BenchParam) * time.Millisecond)
}

func TestPool(t *testing.T) {
	pool, _ := NewPool(10)
	for i := 0; i < 10; i++ {
		i := i
		_ = pool.Execute(context.TODO(), func(ctx context.Context) {
			t.Logf("i:%d", i)
		})
	}
	common.Wait()
}

//	gopool_test.go:66: memory usage:107 MB
//
// --- PASS: TestNoPool (0.48s)
func TestNoPool(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < TestSize; i++ {
		wg.Add(1)
		go func() {
			demoFunc()
			wg.Done()
		}()
	}

	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

//	gopool_test.go:89: memory usage:75 MB
//	gopool_test.go:90: running worker:10000
//	gopool_test.go:91: free worker:0
//
// --- PASS: TestHasPool (1.35s)
func TestHasPool(t *testing.T) {
	pool, _ := NewPool(PoolSize)
	defer pool.Shutdown()
	var wg sync.WaitGroup
	for i := 0; i < TestSize; i++ {
		wg.Add(1)
		_ = pool.Execute(context.TODO(), func(ctx context.Context) {
			demoFunc()
			wg.Done()
		})

	}
	wg.Wait()

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
	t.Logf("running worker:%d", pool.Running())
	t.Logf("free worker:%d", pool.Free())
}
