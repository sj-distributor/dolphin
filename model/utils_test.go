package model

import (
	"testing"
)

func TestGetRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"empty", 0},
		{"short", 8},
		{"medium", 32},
		{"long", 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetRandomString(tt.length)
			if len(result) != tt.length {
				t.Errorf("GetRandomString(%d) returned length %d, want %d", tt.length, len(result), tt.length)
			}
		})
	}

	// 验证两次调用返回不同值（概率极低相同）
	t.Run("unique", func(t *testing.T) {
		a := GetRandomString(32)
		b := GetRandomString(32)
		if a == b {
			t.Error("GetRandomString returned identical values for two calls")
		}
	})
}

func TestRegexpReplace(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		start    string
		end      string
		expected string
	}{
		{"normal", "hello[world]end", `\[`, `\]`, "world"},
		{"not_found", "hello", `\[`, `\]`, ""},
		{"invalid_regex_combined", "hello", `[invalid`, `]`, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// invalid_regex_combined — 当组合的 start+end 正则无效时返回原字符串
			result := RegexpReplace(tt.str, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("RegexpReplace(%q, %q, %q) = %q, want %q", tt.str, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	slice := []string{"a", "b", "c"}

	tests := []struct {
		name     string
		item     string
		expected int
	}{
		{"found_first", "a", 0},
		{"found_middle", "b", 1},
		{"found_last", "c", 2},
		{"not_found", "d", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IndexOf(slice, tt.item)
			if result != tt.expected {
				t.Errorf("IndexOf(%v, %q) = %d, want %d", slice, tt.item, result, tt.expected)
			}
		})
	}
}
