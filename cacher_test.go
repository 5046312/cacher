package cacher

import (
	"fmt"
	"testing"
	"time"
)

// File
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

// Memory

func Test_Memory_Init(t *testing.T) {
	mc := DefaultMemoryCacher()
	fmt.Println(mc)
	time.Sleep(60 * time.Second)
}
