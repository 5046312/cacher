package cacher

import (
	"testing"
)

func TestLoadFile(t *testing.T) {
	fc := DefaultFileCacher()
	fc.Set("a", 123, 3)
}

func TestNewFile(t *testing.T) {
	fc := NewFileCacher("./runtime/cache/")
	fc.Set("abc", 123, 300)
}

func TestFileGet(t *testing.T) {
	fc := DefaultFileCacher()
	value := fc.Get("abc")
	t.Log(value)
}

func TestClear(t *testing.T) {
	fc := NewFileCacher("./runtime/cache/")
	fc.Clear()
}
