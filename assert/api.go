package assert

import (
	"github.com/shura1014/common/goerr"
	"github.com/shura1014/common/utils/objectutil"
	"github.com/shura1014/common/utils/stringutil"
	"reflect"
)

func Panic(option goerr.Option) {
	option.Skip = option.Skip + 1
	panic(goerr.WithOption(option))
}

func Equal(value, expect any, errCode goerr.Code, args ...any) {
	if FilterEmptyAndFunc(value, expect) {
		return
	}
	if value == expect {
		return
	}

	if reflect.DeepEqual(value, expect) {
		return
	}

	panic(goerr.WithOption(goerr.Option{Code: errCode.Code(), Msg: errCode.Message(), Args: append(args), Skip: 1}))
}

// IsZero 如果是零值，抛出异常
func IsZero(value any, errCode goerr.Code, args ...any) {
	if objectutil.IsZero(value) {
		panic(goerr.WithOption(goerr.Option{Code: errCode.Code(), Msg: errCode.Message(), Args: append(args), Skip: 1}))
	}
}

// IsNil 如果为空，抛出异常
func IsNil(value any, errCode goerr.Code, args ...any) {
	if objectutil.IsNil(value) {
		panic(goerr.WithOption(goerr.Option{Code: errCode.Code(), Msg: errCode.Message(), Args: append(args), Skip: 1}))
	}
}

// Error 抛出异常
func Error(errCode goerr.Code, args ...any) {
	panic(goerr.WithOption(goerr.Option{Code: errCode.Code(), Msg: errCode.Message(), Args: append(args), Skip: 1}))
}

// IsNumber 如果不是数字类型，抛出异常
func IsNumber(value string, errCode goerr.Code, args ...any) {
	if !stringutil.IsNumeric(value) {
		panic(goerr.WithOption(goerr.Option{Code: errCode.Code(), Msg: errCode.Message(), Args: append(args), Skip: 1}))
	}
}

// IsTrue 如果是false，抛出异常
func IsTrue(value bool, errCode goerr.Code, args ...any) {
	if !value {
		panic(goerr.WithOption(goerr.Option{Code: errCode.Code(), Msg: errCode.Message(), Args: append(args), Skip: 1}))
	}
}
