package ch07

func twoSum(nums []int, target int) []int {
	indices := make(map[int]int)
	for i, num := range nums {
		if index, ok := indices[num]; ok {
			return []int{index, i}
		}

		indices[target-num] = i
	}

	return []int{-1, -1}
}
