package ch06

func expand(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	if left+1 > right-1 {
		return ""
	}

	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	longest := ""

	for i := 0; i < len(s); i++ {
		sub := expand(s, i, i)
		if len(sub) > len(longest) {
			longest = sub
		}

		sub = expand(s, i, i+1)
		if len(sub) > len(longest) {
			longest = sub
		}
	}

	return longest
}
