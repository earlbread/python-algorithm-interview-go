package ch15

import (
	"container/heap"
	"sort"
)

// Heap
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func findKthLargest(nums []int, k int) int {
	q := &IntMaxHeap{}
	for _, num := range nums {
		heap.Push(q, num)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(q)
	}

	return heap.Pop(q).(int)
}

// Sort
func findKthLargest2(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}
