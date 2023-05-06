package Test01

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
)

func ErrString() {
	fmt.Println("ErrString")
	errors.New("这是一个错误")
}

func EncodingCase() {
	type user struct {
		id   int32
		Name string
		Age  uint8
	}
	u := user{1, "nick", 18}
	bytes, err := json.Marshal(u)
	fmt.Println(bytes, err)
	u1 := user{}
	err = json.Unmarshal(bytes, &u1)
	fmt.Println(u1, err)

	str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(str)
	bytes1, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(bytes1, err)

	str1 := hex.EncodeToString(bytes)
	fmt.Println(str1)

	bytes2, err := hex.DecodeString(str1)
	fmt.Println(bytes2, err)
}
