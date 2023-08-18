package fileutil

import "testing"

func TestListFile(t *testing.T) {
	names := ListFileName(".")
	t.Log(names)
}
