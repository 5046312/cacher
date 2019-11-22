package cacher

import "time"

type File struct {
	PATH string
	EXT  string
}

func (*File) Init(opt map[string]interface{}) {

}
func (*File) Set(key, value interface{}) error {
	return nil
}
func (*File) SetExpire(key, value interface{}, exp time.Duration) error {
	return nil
}
func (*File) Has(key interface{}) bool {
	return false
}
func (*File) Get(key interface{}) (interface{}, error) {
	return nil, nil
}
func (*File) Keys() []interface{} {
	return nil
}
func (*File) Pull(key interface{}) (interface{}, error) {
	return nil, nil
}
func (*File) Remove(key interface{}) bool {
	return false
}
func (*File) Len() int {
	return 0
}
func (*File) Clear() error {
	return nil
}
