package adapter

import (
	"sync"
	"time"
)

const (
	MEMORY_GC_Time = time.Duration(60) * time.Second
)

// Memory Adapter
type MemoryCacher struct {
	sync.Mutex
	Caches  map[string]*CacheItem
	GC_Time time.Duration
}

func DefaultMemoryCacher() *MemoryCacher {
	return NewMemoryCacher(MEMORY_GC_Time)
}

func NewMemoryCacher(gcTime time.Duration) *MemoryCacher {
	return &MemoryCacher{
		Caches:  map[string]*CacheItem{},
		GC_Time: gcTime,
	}
}

func (mc *MemoryCacher) Get(key string) interface{} {
	return nil
}
func (mc *MemoryCacher) Set(key string, val interface{}, exp time.Duration) error {
	return nil
}
func (mc *MemoryCacher) Remove(key string) error {
	return nil
}
func (mc *MemoryCacher) Pull(key string) interface{} {
	return nil

}
func (mc *MemoryCacher) Clear() error {
	return nil

}
