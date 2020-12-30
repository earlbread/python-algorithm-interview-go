package ch09

func isValid(s string) bool {
	op := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]rune, 0)
	for _, c := range s {
		if cp, ok := op[c]; ok {
			stack = append(stack, cp)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
