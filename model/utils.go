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

// 正则截取
func RegexpReplace(str, start string, end string) string {
	reg, _ := regexp.Compile(start + ".+?" + end)
	value := reg.FindString(str)

	reg = regexp.MustCompile(start)
	value = reg.ReplaceAllString(value, "")

	reg = regexp.MustCompile(end)
	value = reg.ReplaceAllString(value, "")
	return value
}
