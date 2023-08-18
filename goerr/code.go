package goerr

type Code interface {
	Code() int
	Message() string
	// SetMessage 占位符动态修改Message
	SetMessage(msg string)
}

type ErrorCode struct {
	ErrCode int    `json:"code"`
	ErrMsg  string `json:"msg"`
}

func (code *ErrorCode) Code() int {
	return code.ErrCode
}

func (code *ErrorCode) Message() string {
	return code.ErrMsg
}

func (code *ErrorCode) SetMessage(msg string) {
	code.ErrMsg = msg
}
