package ch12

func sliceSum(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

func makeSum(result *[][]int, combination, candidates []int, target int) {
	if len(candidates) == 0 {
		return
	}

	sum := sliceSum(combination)

	if sum > target {
		return
	}

	if sum == target {
		*result = append(*result, append([]int{}, combination...))
		return
	}

	for i := 0; i < len(candidates); i++ {
		makeSum(result, append(combination, candidates[i]), candidates[i:], target)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	makeSum(&result, []int{}, candidates, target)
	return result
}
