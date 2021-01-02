package ch23

// Recursive
var cache = map[int]int{
	0: 0,
	1: 1,
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	if result, ok := cache[n]; ok {
		return result
	}

	result := fib(n-2) + fib(n-1)
	cache[n] = result

	return result
}

// Iterative
func fib2(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}

	return b
}
