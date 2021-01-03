package ch23

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	pprev := nums[0]
	prev := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		pprev, prev = prev, max(prev, pprev+nums[i])
	}

	return prev
}
