package adapter

import (
	"sync"
	"time"
)

const (
	MEMORY_GC_Time = time.Duration(10) * time.Second
)

// Memory Adapter
type MemoryCacher struct {
	sync.RWMutex
	Caches  map[string]*CacheItem
	GC_Time time.Duration
}

//
func DefaultMemoryCacher() *MemoryCacher {
	return NewMemoryCacher(MEMORY_GC_Time)
}

//
func NewMemoryCacher(gcTime time.Duration) *MemoryCacher {
	mc := &MemoryCacher{
		Caches:  make(map[string]*CacheItem),
		GC_Time: gcTime,
	}
	go mc.gc()
	return mc
}

// 过期删除，存在则返回
func (mc *MemoryCacher) Get(key string) interface{} {
	mc.RLock()
	defer mc.RUnlock()
	if item, ok := mc.Caches[key]; ok {
		// 存在数据，判断是否过期
		if item.isExpired() {
			// 过期，删除
			mc.Remove(key)
			return nil
		}
		return item.Data
	}
	return nil
}

//
func (mc *MemoryCacher) Set(key string, val interface{}, timeout time.Duration) error {
	mc.Caches[key] = &CacheItem{
		Key:      key,
		Data:     val,
		CreateAt: time.Now(),
		Exp:      timeout * time.Second,
	}
	return nil
}

//
func (mc *MemoryCacher) Remove(key string) error {
	delete(mc.Caches, key)
	return nil
}

//
func (mc *MemoryCacher) Pull(key string) interface{} {
	defer mc.Remove(key)
	return mc.Get(key)
}

//
func (mc *MemoryCacher) Clear() error {
	mc.Caches = make(map[string]*CacheItem)
	return nil
}

// GC
func (mc *MemoryCacher) gc() {
	for {
		<-time.After(mc.GC_Time)
		for key := range mc.Caches {
			mc.Get(key)
		}
	}
}
