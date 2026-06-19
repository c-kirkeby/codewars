package main

import "strings"

func IsPalindrome(str string) bool {
	lower := strings.ToLower(str)
	return lower == reverse(lower)
}

func reverse(str string) (result string) {
	for _, value := range str {
		result = string(value) + result
	}
	return result
}
