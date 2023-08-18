package gopool

import "context"

type Func func(ctx context.Context)

type Task struct {
	ctx      context.Context
	taskFunc Func
}
