package ch21

func sum(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret += num
	}

	return ret
}

func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}

	startIndex := 0
	currentGas := 0

	for i := 0; i < len(gas); i++ {
		currentGas += gas[i] - cost[i]

		if currentGas < 0 {
			currentGas = 0
			startIndex = i + 1
		}
	}

	return startIndex
}
