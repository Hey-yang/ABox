package util

import (
	"testing"
)

// 功能测试
func TestParseSize(t *testing.T) {
	if value := Add(1, 4); value != 3 {
		t.Errorf("1+2 expected be 3, but %d got", value)
	}

	if value := Mul(1, 2); value != 2 {
		t.Errorf("1*2 expected be 2, but %d got", value)
	}
}

// 模糊测试
func FuzzParseSize(f *testing.F) {
	f.Skip("模糊测试")
}

// 基准测试
func BenchmarkParseSize(b *testing.B) {
	b.Skip("基准测试")
}

func TestMain(m *testing.M) {
	m.Run()
}
