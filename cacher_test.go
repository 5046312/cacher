package cacher

import (
	"testing"
)

func TestLoadFile(t *testing.T) {
	fc := DefaultFileCacher()
	fc.Set("a", 123, 3)
}

func TestNewFile(t *testing.T) {
	fc := NewFileCacher("./runtime/cache/", ".cache")
	fc.Set("a", 123, 3)
}
func TestFileGet(t *testing.T) {
	fc := DefaultFileCacher()
	fc.Get("abc")
}
