package ch07

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

			left_height := height[stack[len(stack)-1]]
			depth := h - height[top]
			if left_height < h {
				depth = left_height - height[top]
			}

			distance := i - stack[len(stack)-1] - 1

			result += distance * depth
		}

		stack = append(stack, i)
	}

	return result
}
