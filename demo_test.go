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

	Memory.Remove("c")
	Memory.Remove("b")

	val, err = Memory.Get("b")
	t.Log(val, err)

	mLen := Memory.Len()
	t.Log(mLen)

	Memory.Set(map[string]string{"123": "abc"}, 123)

	keys := Memory.Keys()
	t.Log(keys)
}
