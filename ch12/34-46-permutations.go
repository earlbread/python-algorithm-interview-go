package ch12

func makePermutations(result *[][]int, elements []int, permutation []int) {
	if len(elements) == 0 {
		*result = append(*result, append([]int{}, permutation...))
		return
	}

	for i, e := range elements {
		nextElements := append(append([]int{}, elements[:i]...), elements[i+1:]...)
		nextPermutation := append(append([]int{}, permutation...), e)
		makePermutations(result, nextElements, nextPermutation)
	}
}

func permute(nums []int) [][]int {
	result := [][]int{}
	makePermutations(&result, nums, []int{})

	return result
}
