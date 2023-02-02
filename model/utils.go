package model

// 查找数组并返回下标
func IndexOf(str []interface{}, data interface{}) int {
	for k, v := range str {
		if v == data {
			return k
		}
	}

	return -1
}
