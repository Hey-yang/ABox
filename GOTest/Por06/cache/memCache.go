package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type memCache struct {
	maxMemorySize    int64
	maxMemorySizeStr string
	CurrMemorySize   int64
	values           map[string]*memCacheValue
	//写锁
	locker sync.RWMutex
}

type memCacheValue struct {
	val        any
	expireTime time.Time
	size       int64
}

func NewMemCache() Cache {
	return &memCache{}
}

func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)
	fmt.Println("called set max memory")
	return false
}
func (mc *memCache) Set(key string, val any, expire time.Duration) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		size:       GetValueSize(val),
	}
	mc.del(key)
	err := mc.add(key, v)
	if err == nil {
		//fmt.Println(err)
		panic(err)
		return false
	}
	return true
}
func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.CurrMemorySize -= tmp.size
		delete(mc.values, key)
	}
}
func (mc *memCache) add(key string, val *memCacheValue) error {
	cureMemorySize := mc.CurrMemorySize + val.size
	if cureMemorySize > mc.maxMemorySize {
		return errors.New("超出最大内存限制了")
	}
	mc.CurrMemorySize += val.size
	mc.values[key] = val
	return nil
}

func (mc *memCache) Get(key string) (any, bool) {
	mc.locker.RLocker()
	defer mc.locker.RUnlock()
	tmp, ok := mc.get(key)
	return tmp, ok
}
func (mc *memCache) Del(key string) bool {
	return false
}
func (mc *memCache) Exists(key string) bool {
	return false
}
func (mc *memCache) Flush() bool {
	return false
}
func (mc *memCache) Keys() int64 {
	return 0
}
