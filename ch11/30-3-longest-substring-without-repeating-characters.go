package ch11

func lengthOfLongestSubstring(s string) int {
	start := 0
	longest := 0
	used := map[byte]int{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if _, ok := used[c]; ok && used[c] >= start {
			start = used[c] + 1
		}

		longest = max(longest, i-start+1)
		used[c] = i
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
