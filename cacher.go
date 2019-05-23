package cacher

import (
	"time"

	"github.com/5046312/cacher/adapter"
)

var (
	File     adapter.Adapter
	Memory   adapter.Adapter
	Redis    adapter.Adapter
	Memcache adapter.Adapter
)

// File Cacher
func NewFileCacher(path string) adapter.Adapter {
	return adapter.NewFileCacher(path)
}
func DefaultFileCacher() adapter.Adapter {
	return adapter.DefaultFileCacher()
}

// Memory Cacher
func NewMemoryCacher(gcTime time.Duration) adapter.Adapter {
	return adapter.NewMemoryCacher(gcTime)
}

func DefaultMemoryCacher() adapter.Adapter {
	return adapter.DefaultMemoryCacher()
}

// Redis Cacher
func NewRedisCacher() {

}

// Memcache Cacher
func NewMemcacheCacher(server ...string) adapter.Adapter {
	return adapter.NewMemcacheCacher(server...)
}
