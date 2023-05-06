package util

import (
	"testing"
)

type commStruct struct {
	Group         string
	SizeStr       string
	ExpectSize    int64
	ExpectSizeStr string
}

var commTestData []commStruct

// 功能测试
func TestParseSize(t *testing.T) {

}

// 模糊测试
func FuzzParseSize(f *testing.F) {

}

// 基准测试
func BenchmarkParseSize(b *testing.B) {

}

func TestMain(m *testing.M) {

}

func initCommonData() {
	commTestData = []commStruct{
		{"B", "1b", B, "1B"},
		{"B", "100b", 100 * B, "1B"},
		{"KB", "1kb", KB, "1B"},
		{"KB", "100kb", 100 * KB, "1B"},
		{"MB", "1mb", 100, "1B"},
		{"MB", "100mb", 100 * MB, "1B"},
		{"GB", "1gb", 100, "1B"},
		{"GB", "10gb", 100 * GB, "1B"},
		{"TB", "1tb", 100, "1B"},
		{"TB", "100tb", 100 * TB, "1B"},
		{"PB", "1pb", 100, "1B"},
		{"PB", "100pb", 100 * PB, "1B"},
		{"UNKEOWMB", "1k", 2 * KB, "1B"},
		{"UNKEOWMB", "1g", 2 * KB, "1B"},
	}
}
