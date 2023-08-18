package timeutil

import (
	"testing"
)

func TestParse(t *testing.T) {
	parse := Parse("Y-M-D H:m:s")
	t.Log(parse)
}

func TestNow(t *testing.T) {
	t.Log(Now())
}

func TestNowFormat(t *testing.T) {
	t.Log(NowFormat(Parse("Y/M/D")))
}

func TestMilliSeconds(t *testing.T) {
	t.Log(MilliSeconds())
}

func TestConvert(t *testing.T) {
	t.Log(Convert("1000000"))
}
