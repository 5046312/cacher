package gocacher

import "time"

var Memory cacher = &memoryCacher{
	gc:    time.Minute,
	cache: map[interface{}]*cacherItem{},
}

type memoryCacher struct {
	cache map[interface{}]*cacherItem
	gc    time.Duration
}

func (mc *memoryCacher) Init(config map[string]interface{}) cacher {
	return mc
}

func (mc *memoryCacher) Clone(config map[string]interface{}) cacher {
	newCacher := &memoryCacher{}
	newCacher.Init(config)
	return newCacher
}

func (mc *memoryCacher) Set(key, value interface{}) error {
	ci := &cacherItem{
		key: key,
		val: value,
		exp: nil,
	}
	mc.cache[key] = ci
	return nil
}
func (mc *memoryCacher) SetExpire(key, value interface{}, exp time.Duration) error {
	ci := &cacherItem{
		key: key,
		val: value,
		exp: time.Now().Add(exp),
	}
	// exp 不大于 0 时，为永久缓存
	if exp <= 0 {
		ci.exp = time.Date(3018, 11, 23, 22, 44, 0, 0, time.Local)
	}
	mc.cache[key] = ci
	return nil
}
func (mc *memoryCacher) Has(key interface{}) bool {
	_, exist := mc.cache[key]
	return exist
}
func (mc *memoryCacher) Get(key interface{}) (interface{}, error) {
	val, exist := mc.cache[key]
	if exist && val.expired() {
		// 判断是否过期
		mc.Remove(key)
		return nil, nil
	}
	return val, nil
}
func (mc *memoryCacher) Keys() []interface{} {
	keys := make([]interface{}, 0, mc.Len())
	for k := range mc.cache {
		keys = append(keys, k)
	}
	return keys
}
func (mc *memoryCacher) Pull(key interface{}) (interface{}, error) {
	val, err := mc.Get(key)
	mc.Remove(key)
	return val, err
}
func (mc *memoryCacher) Remove(key interface{}) bool {
	delete(mc.cache, key)
	return true
}
func (mc *memoryCacher) Len() int {
	return len(mc.cache)
}
func (mc *memoryCacher) Clear() error {
	mc.cache = map[interface{}]*cacherItem{}
	return nil
}
