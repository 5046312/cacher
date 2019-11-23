package gocacher

import (
	"time"
)

type cacher interface {
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

var Memory cacher = &memoryCacher{}
