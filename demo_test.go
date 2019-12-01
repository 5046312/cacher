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

func Test_File(t *testing.T) {
	type test struct {
		A string
	}
	err := File.Set("a", &test{"abc"})
	t.Log(err)

	// val, err := File.Get("a")
	// t.Log(val, err)

	// val, err = File.Get("b")
	// t.Log(val, err)

	// err = File.Set("a", 456)
	// t.Log(err)

	// val, err = File.Get("a")
	// t.Log(val, err)

	// File.Remove("c")
	// File.Remove("a")

	// mLen := File.Len()
	// t.Log(mLen)
	// //
	// File.Set("a", false)
	// //
	// keys := File.Keys()
	// t.Log(keys)
	//
	//t.Log(File.Get("a"))
	//
	//File.Clear()
	//t.Log(File.Keys())
	//
	//fc := File.Clone(nil)
	//fc.Set("abc", "123")
	//t.Log(fc.Keys())
	//t.Log(fc.Pull("abc"))
	//t.Log(fc.Get("abc"))
}
