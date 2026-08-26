package utils

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

var uniqueColumnNamePattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

// UniqueRecordExists returns true when another row already uses the same value.
func UniqueRecordExists(db *gorm.DB, model any, fieldName string, value any, scope map[string]any, excludeID string) (bool, error) {
	if db == nil || model == nil || isEmpty(value) {
		return false, nil
	}

	trimmedFieldName := normalizeUniqueColumnName(fieldName)
	if trimmedFieldName == "" || trimmedFieldName == "unknown_field" {
		return false, nil
	}

	query := activeUniqueRecordQuery(db, model).Where(map[string]any{trimmedFieldName: value})
	for scopeField, scopeValue := range scope {
		trimmedScopeField := normalizeUniqueColumnName(scopeField)
		if trimmedScopeField == "" || scopeValue == nil {
			continue
		}
		query = query.Where(map[string]any{trimmedScopeField: scopeValue})
	}

	if trimmedExcludeID := strings.TrimSpace(excludeID); trimmedExcludeID != "" {
		query = query.Where("id <> ?", trimmedExcludeID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func activeUniqueRecordQuery(db *gorm.DB, model any) *gorm.DB {
	return db.Model(model).Where("(is_delete IS NULL OR is_delete = ?)", 1)
}

func normalizeUniqueColumnName(fieldName string) string {
	trimmed := strings.TrimSpace(fieldName)
	if trimmed == "" || trimmed == "unknown_field" {
		return ""
	}
	if !uniqueColumnNamePattern.MatchString(trimmed) {
		return ""
	}
	return strcase.ToSnake(trimmed)
}

func graphQLFieldNameFromColumnName(fieldName string) string {
	trimmed := normalizeUniqueColumnName(fieldName)
	if trimmed == "" || !strings.Contains(trimmed, "_") {
		return trimmed
	}

	parts := strings.Split(trimmed, "_")
	if len(parts) == 0 {
		return trimmed
	}

	var builder strings.Builder
	builder.WriteString(parts[0])
	for _, part := range parts[1:] {
		if part == "" {
			continue
		}
		builder.WriteString(strings.ToUpper(part[:1]))
		if len(part) > 1 {
			builder.WriteString(part[1:])
		}
	}
	return builder.String()
}
