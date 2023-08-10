package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// var goos string = runtime.GOOS
	// fmt.Printf("The operating system is: %s\n", goos)
	// // path := os.Getenv("PATH")
	// // fmt.Printf("Path is %s\n", path)

	// var i1 = 5
	// fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	// var i2 = &i1
	// fmt.Printf("An integer: %d,里面数据的地址: %p, it's location in memory: %p\n", *i2, i2, &i2)

	// arr := [3]int{1, 2, 3}
	// arr[1] = 3
	// fmt.Println(arr)

	// var arr = [5]int{1, 3, 5, 7, 9}
	// // 从数组0下标开始取,一直取到2下标前面一个索引
	// var sce = arr[0:2]
	// fmt.Println(sce) // [1 3]
	// // 切片len = 结束位置 - 开始位置
	// fmt.Println(len(sce)) // 2 - 0 = 2
	// fmt.Println(cap(sce)) // 5 - 0 = 5
	// // 数组地址就是数组首元素的地址
	// fmt.Printf("%p\n", &arr)    // 0xc04200a330
	// fmt.Printf("%p\n", &arr[0]) // 0xc04200a330
	// // 切片地址就是数组中指定的开始元素的地址
	// //  arr[0:2]开始地址为0, 所以就是arr[0]的地址
	// fmt.Printf("%p\n", sce) // 0xc04200a330

	// arr := [5]int{1, 3, 5, 7, 9}
	// sce1 := arr[0:4]
	// sce2 := sce1[0:3]
	// fmt.Println(sce1) // [1 3 5 7]
	// fmt.Println(sce2) // [1 3 5]
	// // 由于底层指向同一数组, 所以修改sce2会影响sce1
	// sce2[1] = 666
	// fmt.Println(sce1) // [1 666 5 7]
	// fmt.Println(sce2) // [1 666 5]
	// fmt.Println(arr)  // [1 666 5 7 9]

	// var dict = map[string]string{"name": "lnj", "age": "33", "gender": "male"}
	// //value, ok := dict["age"]
	// //if(ok){
	// // fmt.Println("有age这个key,值为", value)
	// //}else{
	// // fmt.Println("没有age这个key,值为", value)
	// //}
	// if value, ok := dict["age"]; ok {
	// 	fmt.Println("有age这个key,值为", value)
	// }
	// delete(dict, "name")
	// fmt.Println("dict,值为", dict)
	// for key, value := range dict {
	// 	fmt.Println(key, value)
	// }

	// 注意: 这里不用写type和结构体类型名称
	// var stu2 = struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "lnj",
	// 	age:  33,
	// }
	// fmt.Println(stu2)

	// test2()

	var t time.Time = time.Now()
	// 2006/01/02 15:04:05这个字符串是Go语言规定的, 各个数字都是固定的
	// 字符串中的各个数字可以只有组合, 这样就能按照需求返回格式化好的时间
	str1 := t.Format("2006/01/02 15:04:05")
	fmt.Println(str1)
	str2 := t.Format("2006/01/02")
	fmt.Println(str2)
	str3 := t.Format("15:04:05")
	fmt.Println(str3)
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

func test2() {
	// 如果有异常写在defer中, 并且其它异常写在defer后面, 那么只有defer中的异常会被捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 异常B
		}
	}()

	defer func() {
		panic("异常B")
	}()
	panic("异常A")
}
