package util

import (
	"io"
	"net/http/httptest"
	"os"
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

// 嵌套测试
func TestMul(t *testing.T) {
	t.Errorf("嵌套测试开始...")
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 5) != 6 {
			t.Fatal("expected to get 6,but fail...")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(2, -1) != -6 {
			t.Fatal("expected to get -6,but fail...")
		}
	})
	t.Errorf("嵌套测试结束...")
}

func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if value := Mul(c.A, c.B); value != c.Expected {
				t.Fatalf("%d * %d expected %d,but %d got", c.A, c.B, c.Expected, value)
			}
		})
	}
}

func TestMul3(t *testing.T) {
	caseeMap := make(map[string][]string)
	caseeMap["1"] = []string{"1", "2", "3"}
	caseeMap["2"] = []string{"1", "2", "3"}
	caseeMap["3"] = []string{"1", "2", "3"}
	for s, casee := range caseeMap {
		t.Run(s, func(t *testing.T) {
			if s != "1" {
				t.Fatalf("caseeMap key: %s val :%s,%s,%s", s, casee[0], casee[1], casee[1])
			}
		})
	}
	t.Run("4", func(t *testing.T) {
		string1, ok := caseeMap["3"]
		if ok {
			t.Fatalf("存在值  的 站点是 %s", string1)
		} else {
			t.Fatalf("不存在值 站点不存在")
		}
	})
}

// 单元测试http，如果我们需要测试发送数据或者接收数据，可以采用httptest 测试库进行测试
func TestConn(t *testing.T) {
	req := httptest.NewRequest("GET", "https://www.baidu.com/", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	bytes, _ := io.ReadAll(w.Result().Body)
	if string(bytes) != "hello baidu" {
		t.Fatal("expected hello baidu,but got", string(bytes))
	}
}

// 单元测试钩子 加了下面的代码其他测试用例执行就会执行这个
func TestMain(m *testing.M) {
	Before()
	code := m.Run()
	After()
	os.Exit(code)
}
