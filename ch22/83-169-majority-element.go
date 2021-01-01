package ch22

// Boyer–Moore majority vote algorithm
// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
func majorityElement(nums []int) int {
	count := 1
	major := nums[0]

	for i := 1; i < len(nums); i++ {
		if major == nums[i] {
			count++
		} else if count == 0 {
			count++
			major = nums[i]
		} else {
			count--
		}
	}

	return major
}
