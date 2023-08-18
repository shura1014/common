package goerr

type Option struct {
	Code int
	Skip int
	Data any
	Msg  string
	Args []any
}
