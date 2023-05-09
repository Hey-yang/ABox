package util

import (
	"bytes"
	"testing"
)

/*
然后是基准（压力）测试，这里需要说明一下，压力测试和单元测试最大的区别是单元测试它只执行一次，
成功一次就算成功，而压力测试需要执行多次，任何一次不成功都算失败
*/

// BenchmarkParseSize 基准测试
func BenchmarkParseSize(b *testing.B) {
	b.Skip("基准测试")
}

func BenchmarkAdd(b *testing.B) {
	var n int
	//b.N 是基准测试框架提供的，表示循环的次数，没有具体的限制
	for i := 0; i < b.N; i++ {
		n++
	}
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	//有些测试需要一定的启动和初始化时间，如果从 Benchmark() 函数开始计时会很大程度上影响测试结果的精准性。
	//testing.B 提供了一系列的方法可以方便地控制计时器，从而让计时器只在需要的区间进行测试
	elem := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, v := range elem {
			ret += v
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByByteBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}

// go test -v -bench="." -benchtime=5s benchmark_test.go
//、压力测试自定义测试时间，也就是测试多少秒，默认是测试1秒的压测情况
