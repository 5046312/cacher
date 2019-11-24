package gocacher

import (
	"errors"
	"time"
)

type cacher interface {
	Init(config map[string]interface{}) cacher
	Clone(config map[string]interface{}) cacher
	Set(key string, value interface{}) error
	SetExpire(key string, value interface{}, exp time.Duration) error
	Has(key string) bool
	Get(key string) (interface{}, error)
	Pull(key string) (interface{}, error)
	Remove(key string) bool
	Clear() error
	Keys() []string
	Len() int
}

type cacherItem struct {
	key string
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

	KeyNotExistError = errors.New("key not exist")
)
