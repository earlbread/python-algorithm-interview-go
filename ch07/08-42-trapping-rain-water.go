package ch07

// Stack
func trap(height []int) int {
	result := 0

	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			leftHeight := height[stack[len(stack)-1]]
			depth := h - height[top]
			if leftHeight < h {
				depth = leftHeight - height[top]
			}

			distance := i - stack[len(stack)-1] - 1

			result += distance * depth
		}

		stack = append(stack, i)
	}

	return result
}

// Two Pointer
func trap2(height []int) int {
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
