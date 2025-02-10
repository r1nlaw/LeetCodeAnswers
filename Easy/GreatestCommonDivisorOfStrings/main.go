package main

import (
	"strings"
)

func isDivisor(l int, str1, str2 string, len1, len2 int) bool {
	if len1%l != 0 || len2%l != 0 {
		return false
	}
	f1, f2 := len1/l, len2/l
	return strings.Repeat(str1[:l], f1) == str1 && strings.Repeat(str1[:l], f2) == str2
}

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)

	minStr := min(len1, len2)

	for i := minStr; i > 0; i-- {
		if isDivisor(i, str1, str2, len1, len2) {
			return str1[:i]
		}
	}

	return ""
}
