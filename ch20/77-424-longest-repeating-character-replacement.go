package ch20

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func characterReplacement(s string, k int) int {
	start := 0
	maxCharCount := 0
	charCount := make(map[byte]int)
	longest := 0

	for end := 0; end < len(s); end++ {
		charCount[s[end]]++
		maxCharCount = max(maxCharCount, charCount[s[end]])

		if end-start+1-maxCharCount > k {
			charCount[s[start]]--
			start++
		}

		longest = end - start + 1
	}

	return longest
}
