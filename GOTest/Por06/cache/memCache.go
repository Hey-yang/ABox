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
	//读写锁
	locker sync.RWMutex
	//清除过期缓存时间间隔
	clearExpiredItemInterval time.Duration
}

type memCacheValue struct {
	val        any
	expireTime time.Time
	expire     time.Duration
	size       int64
}

func NewMemCache() Cache {
	mc := &memCache{values: make(map[string]*memCacheValue), clearExpiredItemInterval: time.Second}
	go mc.clearExpiredItem()
	return mc
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
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	tmp, ok := mc.get(key)
	if ok {
		if tmp.expire != 0 && tmp.expireTime.Before(time.Now()) {
			mc.del(key)
			return nil, false
		} else {
			return tmp, true
		}
	}
	return nil, false
}
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.get(key)
	return ok
}
func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue, 0)
	mc.CurrMemorySize = 0
	return true
}
func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))
}

func (mc *memCache) clearExpiredItem() {
	timeTicker := time.NewTicker(mc.clearExpiredItemInterval)
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			for key, item := range mc.values {
				if item.expire != 0 && time.Now().After(item.expireTime) {
					mc.locker.RLock()
					mc.del(key)
					mc.locker.RUnlock()
				}
			}
		}
	}
}
