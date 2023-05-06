package Test

import (
	"Test360.com/Por02/seriver"
	"fmt"
)

// myint32是int32的别名
type myint32 = int32

// 这个就是string的衍生类型
type mystring string

type (
	auyss interface {
		int32 | float64 | ~string
	}
)

func NetTest() {
	fmt.Println("NetTest")
	seriver.TestServer()
}
func NewTest1[T int32 | float64](an T) {
	fmt.Println(an)
}

func NewTest[T auyss](an T) {
	fmt.Println(an)
}
