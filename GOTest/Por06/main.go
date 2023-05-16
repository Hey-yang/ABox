package main

import (
	cache_server "Por06/cache.server"
	"time"
)

func main() {
	//var a = cache.NewMemCache()
	//a.SetMaxMemory("100MB")
	cache := cache_server.NewMemCache()
	cache.SetMaxMemory("100MB")
	cache.Set("int", 1, time.Second)
	cache.Set("bool", true, time.Second)
	cache.Set("data", map[string]any{"a": 1, "a1": 2}, time.Second)
	cache.Set("int", 2)
	cache.Set("bool", false)
	cache.Set("data", map[string]any{"a": 4, "a1": 3})

	cache.Get("int")
	cache.Get("bool")
	cache.Flush()
	cache.Keys()
}
