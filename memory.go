package gocacher

import "time"

var Memory cacher

type memoryCacher struct {
	cache map[string]*cacherItem
	gc    time.Duration
}

func init() {
	cacher := &memoryCacher{
		gc:    gcTime,
		cache: map[string]*cacherItem{},
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
		gc:    mc.gc,
		cache: map[string]*cacherItem{},
	}
	newCacher.runGc(gcTime)
	return newCacher.Init(config)
}

// 设置一个永久时间的缓存
func (mc *memoryCacher) Set(key string, value interface{}) error {
	return mc.SetExpire(key, value, 0)
}
func (mc *memoryCacher) SetExpire(key string, value interface{}, exp time.Duration) error {
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
func (mc *memoryCacher) Has(key string) bool {
	_, exist := mc.cache[key]
	return exist
}
func (mc *memoryCacher) Get(key string) (interface{}, error) {
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
func (mc *memoryCacher) Keys() []string {
	keys := []string{}
	for k, v := range mc.cache {
		if v.expired() {
			mc.Remove(k)
		} else {
			keys = append(keys, k)
		}
	}
	return keys
}
func (mc *memoryCacher) Pull(key string) (interface{}, error) {
	defer mc.Remove(key)
	val, err := mc.Get(key)
	return val, err
}
func (mc *memoryCacher) Remove(key string) bool {
	delete(mc.cache, key)
	return true
}
func (mc *memoryCacher) Len() int {
	return len(mc.cache)
}
func (mc *memoryCacher) Clear() error {
	mc.cache = map[string]*cacherItem{}
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
