package gocacher

import (
	"testing"
)

func Test_Memory(t *testing.T) {
	Memory.Set("a", 123)

	val, err := Memory.Get("a")
	t.Log(val, err)

	val, err = Memory.Get("b")
	t.Log(val, err)
}
