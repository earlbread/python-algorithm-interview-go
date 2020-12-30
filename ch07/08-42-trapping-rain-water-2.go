package ch07

func trap(height []int) int {
	result := 0

	if len(height) <= 2 {
		return 0
	}

	left := 0
	right := len(height) - 1

	leftMax := height[left]
	rightMax := height[right]

	for left < right {
		if height[left] > leftMax {
			leftMax = height[left]
		}

		if height[right] > rightMax {
			rightMax = height[right]
		}

		if leftMax < rightMax {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}

	return result
}
