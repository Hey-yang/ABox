package main

import "Por06/cache"

func main() {
	var a = cache.NewMemCache()
	a.SetMaxMemory("100MB")
}
