package cacher

import "github.com/5046312/cacher/adapter"

var (
	File     adapter.Adapter
	Memory   adapter.Adapter
	Redis    adapter.Adapter
	Memcache adapter.Adapter
)

// File Cacher
func NewFileCacher(path string, ext string) adapter.Adapter {
	return adapter.NewFileCacher(path, ext)
}
func DefaultFileCacher() adapter.Adapter {
	return adapter.DefaultFileCacher()
}

// Memory Cacher
func MemoryCacher() {

}

func RedisCacher() {

}

func MemcacheCacher() {

}
