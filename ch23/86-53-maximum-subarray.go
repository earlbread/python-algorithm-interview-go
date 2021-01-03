package ch23

func maxSubArray(nums []int) int {
	sum := nums[0]
	max := sum

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
