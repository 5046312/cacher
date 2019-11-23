package gocacher

import "time"

var Memory cacher

type memoryCacher struct {
	cache map[interface{}]*cacherItem
	gc    time.Duration
}

func init() {
	cacher := &memoryCacher{
		gc:    gcTime,
		cache: map[interface{}]*cacherItem{},
	}
	cacher.runGc(gcTime)
	Memory = cacher
}

func (mc *memoryCacher) Init(config map[string]interface{}) cacher {
	if val, ok := config["gc"].(time.Duration); ok {
		if val != mc.gc {
			// 修改了gc间隔, 则再起一个新gc协程(旧协程会在到达下次间隔时停止)
			mc.runGc(val)
		}
		mc.gc = val
	}
	return mc
}

// Clone 一个新的缓存对象
func (mc *memoryCacher) Clone(config map[string]interface{}) cacher {
	newCacher := &memoryCacher{
		gc:    gcTime,
		cache: map[interface{}]*cacherItem{},
	}
	newCacher.runGc(gcTime)
	newCacher.Init(config)
	return newCacher
}

// 设置一个永久时间的缓存
func (mc *memoryCacher) Set(key, value interface{}) error {
	return mc.SetExpire(key, value, 0)
}

func (mc *memoryCacher) SetExpire(key, value interface{}, exp time.Duration) error {
	ci := &cacherItem{
		key: key,
		val: value,
		exp: time.Now().Add(exp),
	}
	// exp 不大于 0 时，为永久缓存
	if exp <= 0 {
		ci.exp = farTime
	}
	mc.cache[key] = ci
	return nil
}
func (mc *memoryCacher) Has(key interface{}) bool {
	_, exist := mc.cache[key]
	return exist
}
func (mc *memoryCacher) Get(key interface{}) (interface{}, error) {
	item, exist := mc.cache[key]
	if exist && item.expired() {
		// 判断是否过期
		mc.Remove(key)
		return nil, nil
	} else if !exist {
		return nil, nil
	}
	return item.val, nil
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

// 按执行时的时间规定来执行，如果修改了配置，则停止循环
func (mc *memoryCacher) runGc(gcTime time.Duration) {
	go func() {
		for mc.gc == gcTime {
			for k, v := range mc.cache {
				if v.expired() {
					mc.Remove(k)
				}
			}
			time.Sleep(gcTime)
		}
	}()
}
