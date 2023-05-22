package reflect_test

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func Reflect1() {
	var a float32 = 3.14
	t := reflect.TypeOf(a)  // type:float32
	v := reflect.ValueOf(a) // type:float32
	fmt.Println("类型为:", t, "值为:", v)
	//类型（Type）指type关键字构造很多自定义类型  种类（Kind）指底层的类型
	fmt.Printf("%v %v\n", t.Name(), t.Kind())
	fmt.Println(v.Kind())

	//reflect.ValueOf()返回的是一份值的拷贝，传入要修改的变量的地址，再调用Elem()找到这个指针指向的值，最后使用SetXxx()来修改值。
	// 有时候不一定所有都可修改的，所以可以用CanSet()判断是否可以修改对象。
	t = reflect.TypeOf(&a)  // type:float32
	v = reflect.ValueOf(&a) // type:float32
	v.Elem().SetFloat(1)
	fmt.Println("类型为:", t, "实际类型为:", t.Elem().Kind(), "值为:", a)
	if t.Elem().Kind() == reflect.Float32 {
		fmt.Println("类型为:Float32")
	}
}

func Reflect2() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	v := reflect.ValueOf(stu1)
	fmt.Println("类型为:", t, "底层类型为:", t.Kind(), "值为:", v)
	// reflect.Value与原始值之间可以互相转换。
	// stu2 := reflect.ValueOf(stu1).Interface().(student)
	// fmt.Println(stu2)
}
