package ctx

import (
	"context"
	"testing"
)

func TestCtx(t *testing.T) {
	c := NewValueContext()
	c.SetMap(map[any]any{
		"key1": "value1",
		"key2": "value2",
	})
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	value1 := ctx.Value("key1")
	value2 := ctx.Value("key2")
	value3 := ctx.Value("key3")
	t.Log(value1, value2, value3)

}
