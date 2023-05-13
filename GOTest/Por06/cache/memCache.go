package cache

import "time"

type memCache struct {
	maxMemorySize    int64
	maxMemorySizeStr string
	CurrMemorySize   int64
}

func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	return false
}
func Set(key string, val any, expire time.Duration) bool {

	return false
}
func Get(key string) (any, bool) {
	return nil, false
}
func Del(key string) bool {
	return false
}
func Exists(key string) bool {
	return false
}
func Flush() bool {
	return false
}
func Keys() int64 {
	return 0
}
