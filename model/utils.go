package model

import (
	"math/rand"
	"regexp"
	"time"
)

// 查找数组并返回下标
func IndexOf(str []interface{}, data interface{}) int {
	for k, v := range str {
		if v == data {
			return k
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
