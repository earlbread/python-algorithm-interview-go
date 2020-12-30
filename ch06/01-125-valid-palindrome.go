package ch06

import "strings"

func isAlnum(c byte) bool {
	if (c < 'a' || c > 'z') && (c < '0' || c > '9') {
		return false
	}

	return true
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for left <= right {
		for left <= right && !isAlnum(s[left]) {
			left++
		}

		for left <= right && !isAlnum(s[right]) {
			right--
		}

		if left <= right && s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
