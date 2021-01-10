package ch11

import "container/heap"

type IntMaxHeap [][]int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntMaxHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}

	h := &IntMaxHeap{}
	result := []int{}
	for k, v := range frequency {
		heap.Push(h, []int{v, k})
	}

	for i := 0; i < k; i++ {
		val := heap.Pop(h).([]int)
		result = append(result, val[1])
	}

	return result
}
