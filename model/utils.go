package model

import (
	"math/rand"
	"regexp"
	"time"
)

// IndexOf returns the index of item in slice, or -1 if not found.
func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func GetRandomString(n int) string {
	// 创建一个新的随机数生成器
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = bytes[rng.Intn(len(bytes))]
	}
	return string(result)
}

// RegexpReplace 正则截取：提取 start 和 end 之间的内容
func RegexpReplace(str, start string, end string) string {
	reg, err := regexp.Compile(start + ".+?" + end)
	if err != nil {
		return str
	}
	value := reg.FindString(str)

	regStart, err := regexp.Compile(start)
	if err != nil {
		return str
	}
	value = regStart.ReplaceAllString(value, "")

	regEnd, err := regexp.Compile(end)
	if err != nil {
		return str
	}
	value = regEnd.ReplaceAllString(value, "")
	return value
}
