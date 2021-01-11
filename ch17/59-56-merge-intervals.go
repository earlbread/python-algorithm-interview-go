package ch17

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	for _, interval := range intervals {
		if len(result) > 0 && result[len(result)-1][1] >= interval[0] {
			result[len(result)-1][1] = max(result[len(result)-1][1], interval[1])
		} else {
			result = append(result, interval)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
