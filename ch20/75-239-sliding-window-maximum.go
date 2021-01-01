package ch20

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	q := make([]int, 0)

	for i, num := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if i >= k-1 {
			result = append(result, nums[q[0]])
		}

		if q[0] == i-k+1 {
			q = q[1:]
		}
	}

	return result
}
