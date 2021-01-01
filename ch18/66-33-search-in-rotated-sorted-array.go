package ch18

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	pivot := left

	left = 0
	right = len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		midPivot := (mid + pivot) % len(nums)

		if nums[midPivot] < target {
			left = mid + 1
		} else if nums[midPivot] > target {
			right = mid - 1
		} else {
			return midPivot
		}
	}

	return -1
}
