package ch12

func makeCombinations(result *[][]int, combination []int, start, n, k int) {
	if len(combination) == k {
		*result = append(*result, append([]int{}, combination...))
		return
	}

	for i := start; i <= n; i++ {
		makeCombinations(result, append(combination, i), i+1, n, k)
	}
}

func combine(n int, k int) [][]int {
	result := [][]int{}
	makeCombinations(&result, []int{}, 1, n, k)

	return result
}
