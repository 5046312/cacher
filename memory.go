package gocacher

import "time"

type memoryCacher struct {
	GC time.Duration
}

func (mc *memoryCacher) Clone() *memoryCacher {
	newCacher := &memoryCacher{
		GC: mc.GC,
	}
	return newCacher
}

func (*memoryCacher) Set(key, value interface{}) error {
	return nil
}
func (*memoryCacher) SetExpire(key, value interface{}, exp time.Duration) error {
	return nil
}
func (*memoryCacher) Has(key interface{}) bool {
	return false
}
func (*memoryCacher) Get(key interface{}) (interface{}, error) {
	return nil, nil
}
func (*memoryCacher) Keys() []interface{} {
	return nil
}
func (*memoryCacher) Pull(key interface{}) (interface{}, error) {
	return nil, nil
}
func (*memoryCacher) Remove(key interface{}) bool {
	return false
}
func (*memoryCacher) Len() int {
	return 0
}
func (*memoryCacher) Clear() error {
	return nil
}
