package ch23

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a := 1
	b := 2

	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
