package ch09

func removeDuplicateLetters(s string) string {
	lastOccur := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastOccur[s[i]] = i
	}

	used := make(map[byte]bool)
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if used[c] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > c && lastOccur[stack[len(stack)-1]] > i {
			lastChar := stack[len(stack)-1]
			used[lastChar] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, c)
		used[c] = true
	}

	return string(stack)
}
