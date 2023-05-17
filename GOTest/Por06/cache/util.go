package cache

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	//默认大小为100MB
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit) //转大写
	var BtyeNum int64 = 0
	switch unit {
	case "B":
		BtyeNum = num
	case "KB":
		BtyeNum = num * KB
	case "MB":
		BtyeNum = num * MB
	case "GB":
		BtyeNum = num * GB
	case "TB":
		BtyeNum = num * TB
	case "PB":
		BtyeNum = num * PB
	default:
		num = 0
	}
	if num == 0 {
		log.Println("ParseSize 仅支持 B KB MB GB TB PB")
		num = 100
		BtyeNum = num * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit
	return BtyeNum, sizeStr
}

func GetValueSize(val any) int64 {
	return 0
}
