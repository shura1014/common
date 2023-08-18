package assert

import (
	"github.com/shura1014/common/goerr"
	"reflect"
)

func FilterEmptyAndFunc(expected, actual any) bool {
	if expected == nil && actual == nil {
		return true
	}

	if isFunction(expected) || isFunction(actual) {
		return false
	}
	return false
}

func isFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}

var EqualErr = &goerr.ErrorCode{ErrCode: 500, ErrMsg: "actual %v expect %v"}

func Assert(actual, expect any) {
	Equal(actual, expect, EqualErr, actual, expect)
}
