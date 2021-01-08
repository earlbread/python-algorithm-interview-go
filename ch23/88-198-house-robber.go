package ch23

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	pprev := 0
	prev := 0

	for _, num := range nums {
		pprev, prev = prev, max(prev, pprev+num)
	}

	return prev
}
