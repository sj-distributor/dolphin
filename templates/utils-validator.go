package templates

var Validator = `package utils
import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/duke-git/lancet/v2/convertor"
)

// generateMap ...
func generateMap(arr []string) map[string]interface{} {
	mapData := make(map[string]interface{})
	if len(arr) > 0 {
		for _, v := range arr {
			field := strings.Split(v, ":")
			if len(field) == 2 {
				mapData[field[0]] = field[1]
			}
		}
	}

	return mapData
}

// Len ...
func Len(value string, length interface{}) error {
	maxLength, err := convertor.ToInt(length)
	if err != nil {
		return err
	}

	if len(value) > int(maxLength) {
		return fmt.Errorf(fmt.Sprintf("max length is %d", maxLength))
	}
	return nil
}

// Min ...
func Min[T int64 | float64](value T, min interface{}) error {
	minValue, err := convertor.ToInt(min)
	if err != nil {
		return err
	}

	if int64(value) < minValue {
		return fmt.Errorf("cannot be greater than " + convertor.ToString(min))
	}
	return nil
}

// Max ...
func Max[T int64 | float64](value T, max interface{}) error {
	maxValue, err := convertor.ToInt(max)
	if err != nil {
		return err
	}

	if int64(value) > maxValue {
		return fmt.Errorf("maximum value is " + convertor.ToString(max))
	}
	return nil
}

// CheckStrLen ...
func CheckStrLen(mapData map[string]interface{}, field reflect.Value) error {
	if mapData["len"] != nil {
		value := field.Interface()
		newValue := field.String()
		if field.Kind() == reflect.Ptr && value.(*string) != nil {
			newValue = string(*value.(*string))
		}
		if err := Len(newValue, mapData["len"]); err != nil {
			return err
		}
	}
	return nil
}

// CheckMinAndMax ...
func CheckMinAndMax(mapData map[string]interface{}, field reflect.Value) error {
	value := field.Interface()
	newValue := int64(value.(int64))
	if field.Kind() == reflect.Ptr && value.(*int64) != nil {
		newValue = int64(*value.(*int64))
	}

	if mapData["min"] != nil {
		if err := Min(newValue, mapData["min"]); err != nil {
			return err
		}
	}

	if mapData["max"] != nil {
		if mapData["min"] != nil {
			min, _ := convertor.ToInt(mapData["min"])
			max, _ := convertor.ToInt(mapData["max"])
			if max < min {
				return fmt.Errorf("max cannot be less than min")
			}
		}
		if err := Max(newValue, mapData["max"]); err != nil {
			return err
		}
	}

	return nil
}

// CheckRuleValue ...
func CheckRuleValue(mapData map[string]interface{}, field reflect.Value) error {
	if mapData["type"] != nil {
		kind := field.Kind()
		rl, ok := Rule[mapData["type"].(string)] // Rule 为正则规则
		if !ok || rl["rgx"] == nil {
			return fmt.Errorf("type " + mapData["type"].(string) + " is empty")
		}

		var newValue string
		if kind == reflect.Ptr {
			if !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.String {
					newValue = elem.String()
				} else if elem.Kind() == reflect.Int || elem.Kind() == reflect.Int64 {
					newValue = fmt.Sprintf("%d", elem.Int())
				}
			}
		} else if kind == reflect.String {
			newValue = field.String()
		} else if kind == reflect.Int || kind == reflect.Int64 {
			newValue = fmt.Sprintf("%d", field.Int())
		}

		if !regexp.MustCompile(rl["rgx"].(string)).MatchString(newValue) {
			return fmt.Errorf(rl["msg"].(string))
		}
	}
	return nil
}

// Struct ...
func Struct(field reflect.Value, tag reflect.StructTag) error {
	validator := tag.Get("validator")

	if validator != "" {
		arr := strings.Split(validator, ";")
		mapData := generateMap(arr)

		value := field.Interface()
		title := tag.Get("json")

		switch reflect.TypeOf(value).String() {
		// 字符串长度限制
		case "string", "*string":
			if err := CheckStrLen(mapData, field); err != nil {
				return fmt.Errorf(title + " " + err.Error())
			}

		// 整数最小值、最大值限制
		case "int64", "*int64":
			if err := CheckMinAndMax(mapData, field); err != nil {
				return fmt.Errorf(title + " " + err.Error())
			}
		}

		// 不允许修改字段
		if mapData["edit"] != nil && mapData["edit"] == "no" {
			return fmt.Errorf(title + " editing not allowed")
		}

		// 正则验证
		if err := CheckRuleValue(mapData, field); err != nil {
			return fmt.Errorf(title + " " + err.Error())
		}
	}

	return nil
}

// Validate ...
func Validate(item interface{}) error {
	data := reflect.ValueOf(item)
	elem := data.Elem()
	elemKey := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		if strings.Contains(elemKey.Field(i).Type.String(), "*") {
			continue
		}
		tag := elemKey.Field(i).Tag
		if err := Struct(elem.Field(i), tag); err != nil {
			return err
		}
	}
	return nil
}
`
