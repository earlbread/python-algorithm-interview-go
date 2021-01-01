package ch22

import "strconv"

func calculate(left_result []int, right_result []int, op rune) []int {
	result := make([]int, 0)
	for _, left := range left_result {
		for _, right := range right_result {
			val := 0
			if op == '+' {
				val = left + right
			} else if op == '-' {
				val = left - right
			} else {
				val = left * right
			}

			result = append(result, val)
		}
	}

	return result
}

func diffWaysToCompute(input string) []int {
	result := make([]int, 0)

	for i, c := range input {
		if c == '-' || c == '+' || c == '*' {
			left_result := diffWaysToCompute(input[:i])
			right_result := diffWaysToCompute(input[i+1:])
			op := c

			result = append(result, calculate(left_result, right_result, op)...)
		}
	}

	if len(result) == 0 {
		num, _ := strconv.Atoi(input)
		result = append(result, num)
	}

	return result
}
