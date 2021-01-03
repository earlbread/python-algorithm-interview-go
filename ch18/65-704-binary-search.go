package ch18

// Iterative
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// Recursive
func binarySearch(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	result := -1
	if nums[mid] < target {
		result = binarySearch(nums, target, mid+1, right)
	} else if nums[mid] > target {
		result = binarySearch(nums, target, left, mid-1)
	} else {
		result = mid
	}

	return result
}

func search2(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}
