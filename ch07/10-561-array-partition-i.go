package ch07

import "sort"

func arrayPairSum(nums []int) int {
	result := 0

	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}

	return result
}
