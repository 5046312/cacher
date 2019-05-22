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
	sync.RWMutex
	Caches  map[string]*CacheItem
	GC_Time time.Duration
}

func DefaultMemoryCacher() *MemoryCacher {
	return NewMemoryCacher(MEMORY_GC_Time)
}

func NewMemoryCacher(gcTime time.Duration) *MemoryCacher {
	return &MemoryCacher{
		Caches:  make(map[string]*CacheItem),
		GC_Time: gcTime,
	}
}

// 检测key是否过期
// 过期则删除，返回 true
// 未过期，返回false
func (mc *MemoryCacher) isExp(item *CacheItem) bool {
	if item.isExpired() {
		return true
	}
	return false
}

//
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
	val := mc.Get(key)
	defer mc.Remove(key)
	return val

}
func (mc *MemoryCacher) Clear() error {
	mc.Caches = make(map[string]*CacheItem)
	return nil
}
