package ch09

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(T))

	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			last := stack[len(stack)-1]
			result[last] = i - last

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}
