package ch21

func leastInterval(tasks []byte, n int) int {
	counts := map[byte]int{}
	maxCount := 0

	for _, task := range tasks {
		counts[task]++
		maxCount = max(maxCount, counts[task])
	}

	time := (maxCount - 1) * (n + 1)

	for _, v := range counts {
		if maxCount == v {
			time++
		}
	}

	return max(len(tasks), time)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
