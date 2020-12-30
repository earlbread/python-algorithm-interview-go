package ch07

func productExceptSelf(nums []int) []int {
	result := make([]int, 0)

	prev := 1
	for i := 0; i < len(nums); i++ {
		result = append(result, prev)
		prev *= nums[i]
	}

	prev = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= prev
		prev *= nums[i]
	}

	return result
}
