package util

import (
	"fmt"
	"net/http"
)

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func Before() {
	fmt.Println("Before all tests...")
}

func After() {
	fmt.Println("After all tests...")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello baidu"))
}

func Hello() string {
	return "Hello World"
}
