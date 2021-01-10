package ch11

func numJewelsInStones(J string, S string) int {
	jewelSet := map[rune]bool{}
	for _, r := range J {
		jewelSet[r] = true
	}

	count := 0
	for _, r := range S {
		if jewelSet[r] {
			count++
		}
	}

	return count
}
