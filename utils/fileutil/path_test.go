package fileutil

import (
	"testing"
)

func TestListFile(t *testing.T) {
	names := ListFileName(".")
	t.Log(names)
}

func TestDir(t *testing.T) {
	dir := Dir("./etc")
	t.Log(dir)
}

func TestCallDir(t *testing.T) {
	t.Log(callDir())
}

func callDir() string {
	return CallDir(".")
}
