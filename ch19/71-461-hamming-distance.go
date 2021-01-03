package ch19

func hammingDistance(x int, y int) int {
	xor := x ^ y
	distance := 0
	for xor != 0 {
		xor = xor & (xor - 1)
		distance++
	}

	return distance
}
