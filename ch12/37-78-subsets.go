package ch12

func makeSubsets(result *[][]int, subset, nums []int) {
	*result = append(*result, append([]int{}, subset...))

	for i := 0; i < len(nums); i++ {
		makeSubsets(result, append(subset, nums[i]), nums[i+1:])
	}
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	makeSubsets(&result, []int{}, nums)
	return result
}
