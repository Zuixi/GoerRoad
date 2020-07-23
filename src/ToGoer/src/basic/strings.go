package basic

import (
	"strings"
)

func Compare(a, b string) int {
	return strings.Compare(a, b)
}

func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// reports whether any unicode code points in chars are within s
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// 判断在Unicode大小写折叠下s和t（解释为UTF-8字符串）是否相等，这是不区分大小写的更通用形式
func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}