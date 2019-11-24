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

	Memory.Set("a", false)

	keys := Memory.Keys()
	t.Log(keys)

	t.Log(Memory.Get("a"))

	Memory.Clear()
	t.Log(Memory.Keys())

	mc := Memory.Clone(nil)
	mc.Set("abc", "123")
	t.Log(mc.Keys())
	t.Log(mc.Pull("abc"))
	t.Log(mc.Get("abc"))
}
