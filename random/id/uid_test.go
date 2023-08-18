package id

import (
	"math"
	"os"
	"strconv"
	"testing"
	"time"
)

// uid_test.go:24: rurg5p0pm0cuvk0befki68y100qlujzd
// uid_test.go:24: rurg5p0pm0cuvk0befm06o200a5vq5cm
// uid_test.go:24: rurg5p0pm0cuvk0befm5l4300e4bst3p
// uid_test.go:24: rurg5p0pm0cuvk0befma7s400i57r4lg
// uid_test.go:24: rurg5p0pm0cuvk0befme2o500y5jlob7
// uid_test.go:24: rurg5p0pm0cuvk0befmh5s600c8fb4l0
// uid_test.go:24: rurg5p0pm0cuvk0befmv1s700inle2ld
// uid_test.go:24: rurg5p0pm0cuvk0befmy4w800db8erxu
// uid_test.go:24: rurg5p0pm0cuvk0befn1zs900fruedbi
// uid_test.go:24: rurg5p0pm0cuvk0befnz68a00qfzcogx
func TestStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := Str()
		t.Log(str)
	}
}

func TestTime(t *testing.T) {
	nano := time.Now().UnixNano()
	t.Log(nano)
	t.Log(int64(math.Pow(32, 12)))
	t.Log(strconv.FormatInt(nano, 36))

	date := time.Date(2023, time.August, 1, 0, 0, 0, 0, time.Local)
	since := time.Since(date).Nanoseconds()
	t.Log(since)
	t.Log(strconv.FormatInt(since, 36))

}

func TestGetMacStr(t *testing.T) {
	t.Log(GetMacStr())
	t.Log(math.Pow(36, 4))
}

func TestPid(t *testing.T) {
	t.Log(os.Getpid())
	s := strconv.FormatInt(int64(os.Getpid()), 36)
	t.Log(s)
}
