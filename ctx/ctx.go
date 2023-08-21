package ctx

import (
	"context"
	"github.com/shura1014/common/container/concurrent"
	"github.com/shura1014/common/utils/stringutil"
)

type MapValueCtx struct {
	context.Context
	data *concurrent.AnyMap[any]
}

func (c *MapValueCtx) Value(key any) any {
	return c.data.Get(key)
}

func (c *MapValueCtx) SetValue(k, v any) {
	c.data.Put(k, v)
}

func (c *MapValueCtx) SetMap(data map[any]any) {
	c.data.PutAll(data)
}

func (c *MapValueCtx) Remove(keys ...any) {
	c.data.Remove(keys...)
}

func (c *MapValueCtx) String() string {
	return stringutil.ToString(c.Context) + ".MapValueCtx: " + c.data.String()
}

func NewValueContext(parent ...context.Context) *MapValueCtx {
	var pc context.Context
	if len(parent) > 0 {
		pc = parent[0]
	} else {
		pc = context.Background()
	}
	return &MapValueCtx{pc, concurrent.NewAnyMap[any]()}
}
