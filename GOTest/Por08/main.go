package main

import (
	"fmt"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	// path := os.Getenv("PATH")
	// fmt.Printf("Path is %s\n", path)

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var i2 = &i1
	fmt.Printf("An integer: %d,里面数据的地址: %p, it's location in memory: %p\n", *i2, i2, &i2)
}

func init() {
	// fmt.Println("init 方法执行")
	// var s, s1 = testFunc()
	// fmt.Println("init 方法执行2", s, s1)
	var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}

	fmt.Println("init 方法执行:", prompt)
}

func testFunc() (int32, int32) {
	return 0, 2
}
