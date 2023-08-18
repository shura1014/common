package env

import (
	"os"
	"testing"
)

// map[a:b e: f:g h:i j:k m:]
func TestInitArgs(t *testing.T) {
	_ = os.Setenv("a", "a")
	_ = os.Setenv("c", "d")
	args := []string{"--a", "b", "c", "d", "--e", "-f", "g", "-h=i", "--j=k", "l", "-m"}
	InitArgs(args)

	t.Log(GetEnv("a")) // expect b
	t.Log(GetEnv("c")) // expect d
	t.Log(GetAll())
}
