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

// 判断interface是否为nil
func IsNil(i interface{}) bool {
  if IsEmpty(i) {
    return true
  }

  tye := reflect.TypeOf(i).String()
  if tye == "int" && int(i.(int)) == 0 || tye == "int64" && int64(i.(int64)) == 0 {
    return true
  } else if tye == "*int" && int(*i.(*int)) == 0 && tye == "*int64" && int64(*i.(*int64)) == 0 {
    return true
  }

  vi := reflect.ValueOf(i)
  if vi.Kind() == reflect.Ptr {
    return vi.IsNil()
  }
  return false
}

// stringToArray
func StrToArr(data, split string) []string {
	if IsEmpty(data) {
		return nil
	}
	return strings.Split(data, split)
}

// StructToMap
func StructToMap(obj interface{}) map[string]interface{} {
  t := reflect.TypeOf(obj)
  v := reflect.ValueOf(obj)

  var data = make(map[string]interface{})
  for i := 0; i < t.NumField(); i++ {
    name := t.Field(i).Name
    data[strcase.ToLowerCamel(name)] = v.Field(i).Interface()
  }
  return data
}

// 数组差集
func Difference[T string](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range b {
		m[item] = true
	}

	var result []T
	for _, item := range a {
		if _, ok := m[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}
`
