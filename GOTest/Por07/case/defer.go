package _case

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func DeferCase1() {
	fmt.Println("开始1")
	defer func() {
		fmt.Println("结束  匿名1")
	}()
	defer f1()

	defer f2()()
	fmt.Println("开始2")

}

func DeferCase2() {
	i := 1
	defer func(j int) {
		fmt.Println("defer j ", j)
	}(i + 1)

	defer func() {
		fmt.Println("defer j ", i)
	}()
	i = 99
	fmt.Println("i =  ", i)
}

func FileReadCase() {
	file, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("go.mod 的资源释放")
		file.Close()
	}()
	var data []byte
	buf := make([]byte, 1024)
	t1 := time.Now()
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}
	// 将data切片转为字符串即使文件内容
	fmt.Println(string(data))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}

func f1() {
	fmt.Println("结束  具名2")
}

func f2() func() {
	fmt.Println("结束  具名3")
	return func() {
		fmt.Println("结束  具名3的返回函数")
	}

}
