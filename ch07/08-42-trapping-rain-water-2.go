package ch07

func trap(height []int) int {
	result := 0

	if len(height) <= 2 {
		return 0
	}

	left := 0
	right := len(height) - 1

	left_max := height[left]
	right_max := height[right]

	for left < right {
		if height[left] > left_max {
			left_max = height[left]
		}

		if height[right] > right_max {
			right_max = height[right]
		}

		if left_max < right_max {
			result += left_max - height[left]
			left++
		} else {
			result += right_max - height[right]
			right--
		}
	}

	return result
}
