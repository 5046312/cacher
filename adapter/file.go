package adapter

import "time"

type CacherFile struct {
}

func SetFile(fc *CacherFile) {

}

func (CacherFile) Get(key string) interface{} {
	panic("implement me")
}

func (CacherFile) All(keys []string) []interface{} {
	panic("implement me")
}

func (CacherFile) Set(key string, val interface{}, timeout time.Duration) error {
	panic("implement me")
}

func (CacherFile) Inc(key string) error {
	panic("implement me")
}

func (CacherFile) Dec(key string) error {
	panic("implement me")
}

func (CacherFile) Remove(key string) error {
	panic("implement me")
}

func (CacherFile) Pull(key string) interface{} {
	panic("implement me")
}

func (CacherFile) IsExist(key string) bool {
	panic("implement me")
}

func (CacherFile) Clear() error {
	panic("implement me")
}

func (CacherFile) Tag(tag string) error {
	panic("implement me")
}

func DefaultCacherFileConfig() *CacherFile {
	return &CacherFile{}
}
