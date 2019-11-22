package cacher

import "time"

type Memory struct {
	GCTime time.Duration
}

func (*Memory) Init(opt map[string]interface{}) {

}
func (*Memory) Set(key, value interface{}) error {
	return nil
}
func (*Memory) SetExpire(key, value interface{}, exp time.Duration) error {
	return nil
}
func (*Memory) Has(key interface{}) bool {
	return false
}
func (*Memory) Get(key interface{}) (interface{}, error) {
	return nil, nil
}
func (*Memory) Keys() []interface{} {
	return nil
}
func (*Memory) Pull(key interface{}) (interface{}, error) {
	return nil, nil
}
func (*Memory) Remove(key interface{}) bool {
	return false
}
func (*Memory) Len() int {
	return 0
}
func (*Memory) Clear() error {
	return nil
}
