package ch20

func minWindow(s string, t string) string {
	missing := len(t)

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	result := ""
	left := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		if need[c] > 0 {
			missing--
		}

		need[c]--

		if missing == 0 {
			for left < right && need[s[left]] < 0 {
				need[s[left]]++
				left++
			}

			if result == "" || len(s[left:right+1]) < len(result) {
				result = s[left : right+1]
			}

			need[s[left]]++
			missing++
			left++
		}
	}

	return result
}
