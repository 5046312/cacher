package cacher

import "testing"

func TestLoadFile(t *testing.T) {
	fc := FileConfig()
	ca := SetFile(fc)
	ca.Get("asd")
}
