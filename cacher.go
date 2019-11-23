package gocacher

import (
	"time"
)

type cacher interface {
	Init(config map[string]interface{}) cacher
	Clone(config map[string]interface{}) cacher
	Set(key, value interface{}) error
	SetExpire(key, value interface{}, exp time.Duration) error
	Has(key interface{}) bool
	Get(key interface{}) (interface{}, error)
	Pull(key interface{}) (interface{}, error)
	Remove(key interface{}) bool
	Clear() error
	Keys() []interface{}
	Len() int
}

type cacherItem struct {
	key interface{}
	val interface{}
	exp time.Time
}

// 该数据是否过期
func (ci *cacherItem) expired() bool {
	return ci.exp.Before(time.Now())
}

var (
	gcTime  = time.Hour
	farTime = time.Date(3018, 11, 23, 22, 44, 0, 0, time.Local)
)
