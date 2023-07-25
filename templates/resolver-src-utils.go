package templates

var ResolverSrcUtils = `package utils

import (
	"os"
	"strings"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return true
	}
	return false
}

func EnsureDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	switch v.(type) {
	case string:
		if v == "" {
			return true
		}
	default:
		return false
	}
	return false
}

// 查找数组并返回下标
func StrIndexOf(str []string, data interface{}) int {
	for k, v := range str {
		if v == data {
			return k
		}
	}

	return -1
}

// stringToArray
func StrToArr(data, split string) []string {
	if IsEmpty(data) {
		return nil
	}
	return strings.Split(data, split)
}`
