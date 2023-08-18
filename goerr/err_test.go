package goerr

import (
	"errors"
	"github.com/shura1014/common/assert"
	"testing"
)

func TestWrap(t *testing.T) {
	e1 := Text("%s is not nil", "user")
	e2 := Wrap(errors.New("failed"))
	t.Log(e1)
	t.Log(e2)
}

func TestAssert(t *testing.T) {
	assert.IsTrue(1 == 2, &ErrorCode{5000, "%s is not nil"}, "data")
}
