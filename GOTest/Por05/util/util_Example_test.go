package util

import "fmt"

// 实例测试，这个测试有点简单，写个 //Output注释，然后后面是期望输出的结果，如果是一致的，测试通过，否则测试失败
// Output 需要和上面打印的内容一致，否则测试失败
func ExampleHello() {
	fmt.Println(Hello())
	// Output: Hello World
}
