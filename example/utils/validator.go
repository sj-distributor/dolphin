package utils

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/duke-git/lancet/v2/convertor"
	"gorm.io/gorm"
)

// 获取字段名
func GetFieldName(obj any, value any) string {
	if objMap, ok := obj.(map[string]interface{}); ok {
		for key, val := range objMap {
			if val == value {
				return key
			}
		}
	}
	return "unknown_field"
}

// 通用校验方法
func ValidateField(ctx context.Context, obj any, fieldName string, value any, required *string, immutable *string, typeArg *string, minLength *int, maxLength *int, minValue *int, maxValue *int, unique *string, uniqueScope *string) error {

	fieldContext := graphql.GetFieldContext(ctx)
	if !strings.Contains(fieldContext.Field.Name, "create") && immutable != nil && *immutable == "true" {
		return fmt.Errorf("%s cannot be modified", fieldName)
	}

	// 必填校验
	if required != nil && *required == "true" && isEmpty(value) {
		return fmt.Errorf("%s is required", fieldName)
	}

	// 类型校验（正则匹配）
	if typeArg != nil {
		if err := validateType(fieldName, value, *typeArg); err != nil {
			return err
		}
	}

	// 字符串长度校验
	if minLength != nil || maxLength != nil {
		if err := validateStringLength(fieldName, value, minLength, maxLength); err != nil {
			return err
		}
	}

	// 数值范围校验
	if minValue != nil || maxValue != nil {
		if err := validateNumberRange(fieldName, value, minValue, maxValue); err != nil {
			return err
		}
	}

	// 唯一性校验
	if unique != nil && *unique == "true" {
		if err := validateUnique(ctx, obj, fieldName, value, uniqueScope); err != nil {
			return err
		}
	}

	return nil
}

// 判断值是否为空
func isEmpty(value any) bool {
	if value == nil {
		return true
	}
	if str, ok := value.(string); ok && len(str) == 0 {
		return true
	}
	return false
}

// 正则校验
func validateType(fieldName string, value any, typeArg string) error {
	rule, ok := Rule[typeArg]
	if !ok || rule["rgx"] == nil {
		return fmt.Errorf("type %s is empty", typeArg)
	}

	regexPattern, _ := rule["rgx"].(string)
	if strValue, ok := value.(string); ok {
		if !regexp.MustCompile(regexPattern).MatchString(strValue) {
			return fmt.Errorf("%s format is invalid: %s", fieldName, rule["msg"])
		}
	}
	return nil
}

// 字符串长度校验
func validateStringLength(fieldName string, value any, minLength *int, maxLength *int) error {
	strValue, ok := value.(string)
	if !ok {
		return nil // 非字符串不做长度校验
	}

	if minLength != nil && len(strValue) < *minLength {
		return fmt.Errorf("%s must be at least %d characters long", fieldName, *minLength)
	}

	if maxLength != nil && len(strValue) > *maxLength {
		return fmt.Errorf("%s must be at most %d characters long", fieldName, *maxLength)
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
		return fmt.Errorf("must be at least %s", convertor.ToString(min))
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
		return fmt.Errorf("must be at most %s", convertor.ToString(max))
	}
	return nil
}

// 数值范围校验
func validateNumberRange(fieldName string, value any, minValue *int, maxValue *int) error {
	intValue := int64(value.(int64))
	if minValue != nil {
		if err := Min(intValue, *minValue); err != nil {
			return err
		}
	}

	if maxValue != nil {
		if minValue != nil {
			min, _ := convertor.ToInt(minValue)
			max, _ := convertor.ToInt(maxValue)
			if max < min {
				return fmt.Errorf("%s max cannot be less than min", fieldName)
			}
		}

		if err := Max(intValue, *maxValue); err != nil {
			return err
		}
	}

	return nil
}

// validateUnique 唯一性校验
// 通过查询数据库确认该字段值是否已存在
// uniqueScope 为可选条件字段，如 "uid"，表示在同一 uid 下唯一
func validateUnique(ctx context.Context, obj any, fieldName string, value any, uniqueScope *string) error {
	if isEmpty(value) {
		return nil // 空值不做唯一性校验
	}

	// 从 context 获取 db 和 tableName
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok || db == nil {
		return nil // 无数据库连接时跳过校验
	}

	tableName, ok := ctx.Value("tableName").(string)
	if !ok || tableName == "" {
		return nil // 无表名时跳过校验
	}

	query := db.Table(tableName).Where(fieldName+" = ?", value)

	// 可选的唯一性条件范围字段
	if uniqueScope != nil && *uniqueScope != "" {
		if objMap, ok := obj.(map[string]interface{}); ok {
			if scopeValue, exists := objMap[*uniqueScope]; exists && scopeValue != nil {
				query = query.Where(*uniqueScope+" = ?", scopeValue)
			}
		}
	}

	// 排除软删除的记录
	query = query.Where("is_delete = ?", 1)

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil // 查询出错时不阻塞正常请求
	}

	if count > 0 {
		return fmt.Errorf("%s already exists", fieldName)
	}

	return nil
}
