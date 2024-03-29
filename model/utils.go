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
	rand.Seed(time.Now().Unix())
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
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
