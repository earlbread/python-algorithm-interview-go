package ch18

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)

	for _, num := range nums1 {
		set[num] = true
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		if set[num] {
			set[num] = false
			result = append(result, num)
		}
	}

	return result
}
