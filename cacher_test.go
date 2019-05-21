package cacher

import "testing"

func TestLoadFile(t *testing.T) {
	fc := DefaultFileCacher()
	fc.Set("a", 123, 3)
}
