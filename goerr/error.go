package goerr

import (
	"encoding/json"
	"fmt"
	"github.com/shura1014/common/utils/stringutil"
	"runtime"
	"strconv"
	"strings"
)

var defaultSkip = 3

type IStack interface {
	Stack() string
}

func IsStack(err any) bool {
	_, ok := err.(IStack)
	return ok
}

type BizError struct {
	Code
	stack string
	err   error
	Data  any `json:"data"`
}

func (e *BizError) Error() string {
	return fmt.Sprintf("%d:%s", e.Code.Code(), e.Message())
}

func (e *BizError) Json() string {
	builder := strings.Builder{}
	builder.WriteByte('{')
	builder.WriteString("\"code\":")
	builder.WriteString(strconv.Itoa(e.Code.Code()))
	builder.WriteByte(',')
	builder.WriteString("\"msg\":")
	builder.WriteString(fmt.Sprintf("\"%s\"", e.Message()))
	data, err := json.Marshal(e.Data)
	if err == nil {
		builder.WriteByte(',')
		builder.WriteString("\"data\":")
		builder.Write(data)
	}

	builder.WriteByte('}')
	return builder.String()
}

func (e *BizError) DetailMsg() string {
	return e.stack
}

func (e *BizError) Stack() string {
	return e.stack
}

func (e *BizError) String() string {
	return e.Error() + ":" + e.DetailMsg()
}

func Wrapf(err any, msg string, args ...any) *BizError {
	message := msg
	if len(args) > 0 {
		message = fmt.Sprintf(msg, args...)
	}
	switch err.(type) {
	case *BizError:
		return err.(*BizError)
	case error:
		return &BizError{
			Code: &ErrorCode{
				ErrCode: 500,
				ErrMsg:  message,
			},
			stack: DetailMsg(err, defaultSkip),
			err:   err.(error),
		}
	default:
		return &BizError{
			Code: &ErrorCode{
				ErrCode: 500,
				ErrMsg:  message,
			},
			stack: DetailMsg(err, defaultSkip),
		}
	}
}

// Wrap 包装异常
// err 可以是任意类型
// n 默认跳过七层，如果刚好报错的地方包装，那么是不需要传该参数，但是业务经过好几层传递才调用的该函数，需要业务自己计算一下经过了多少层
func Wrap(err any, n ...int) *BizError {
	skip := defaultSkip + 1
	if len(n) > 0 {
		skip += n[0]
	}
	switch err.(type) {
	case *BizError:
		return err.(*BizError)
	//case runtime.Error: /*忽略GoRoot 这一步就不需要了*/
	//	// 忽略
	//	// runtime/panic.go:884
	//	// runtime/panic.go:260
	//	// runtime/signal_unix.go:835
	//	skip += 3
	//	return &BizError{
	//		Code: &ErrorCode{
	//			ErrCode: 500,
	//			ErrMsg:  stringutil.ToString(err),
	//		},
	//		stack: DetailMsg(err, skip),
	//		err:   err.(error),
	//	}
	case error:
		return &BizError{
			Code: &ErrorCode{
				ErrCode: 500,
				ErrMsg:  stringutil.ToString(err),
			},
			stack: DetailMsg(err, skip),
			err:   err.(error),
		}
	default:
		return &BizError{
			Code: &ErrorCode{
				ErrCode: 500,
				ErrMsg:  stringutil.ToString(err),
			},
			stack: DetailMsg(err, skip),
		}
	}
}

func TextSkip(skip int, msg string, args ...any) *BizError {
	skip += defaultSkip
	message := msg
	if len(args) > 0 {
		message = fmt.Sprintf(msg, args...)
	}
	return &BizError{
		Code: &ErrorCode{
			ErrCode: 500,
			ErrMsg:  message,
		},
		stack: DetailMsg(message, skip),
	}
}

func Text(msg string, args ...any) *BizError {
	message := msg
	if len(args) > 0 {
		message = fmt.Sprintf(msg, args...)
	}
	return &BizError{
		Code: &ErrorCode{
			ErrCode: 500,
			ErrMsg:  message,
		},
	}
}

func WithCodeAndData(code Code, data any, args ...any) *BizError {
	if len(args) > 0 {
		code.SetMessage(fmt.Sprintf(code.Message(), args...))
	}
	return &BizError{
		Code:  code,
		Data:  data,
		stack: DetailMsg(code.Message(), defaultSkip),
	}
}

func WithCode(code Code, args ...any) *BizError {

	if len(args) > 0 {
		code.SetMessage(fmt.Sprintf(code.Message(), args...))
	}
	return &BizError{
		Code:  code,
		stack: DetailMsg(code.Message(), defaultSkip),
	}
}

func WithOption(option Option) *BizError {
	msg := option.Msg
	if option.Args != nil {
		msg = fmt.Sprintf(msg, option.Args...)
	}
	return &BizError{
		Code: &ErrorCode{
			ErrCode: option.Code,
			ErrMsg:  msg,
		},
		Data:  option.Data,
		err:   nil,
		stack: DetailMsg(msg, option.Skip+defaultSkip),
	}
}

func NewError(code int, data any, format string, args ...any) *BizError {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}
	return &BizError{
		Code: &ErrorCode{
			ErrCode: code,
			ErrMsg:  msg,
		},
		Data:  data,
		err:   nil,
		stack: DetailMsg(msg, defaultSkip),
	}
}

var (
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}

func DetailMsg(err any, skip int) string {
	var pcs [64]uintptr
	n := runtime.Callers(skip, pcs[:])
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%v", err))
	for _, pc := range pcs[:n] {
		if fn := runtime.FuncForPC(pc - 1); fn != nil {
			file, line := fn.FileLine(pc - 1)
			// 忽略GoRoot
			if goRootForFilter != "" &&
				len(file) >= len(goRootForFilter) &&
				file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			sb.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
		}

	}
	return sb.String()
}

func (e *BizError) Unwrap() error {
	return e.err
}
