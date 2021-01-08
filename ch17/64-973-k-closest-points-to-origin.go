package ch17

import "container/heap"

func dist(point []int) int {
	x := point[0]
	y := point[1]
	return x*x + y*y
}

type PointHeap [][]int

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return dist(h[i]) < dist(h[j]) }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *PointHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func kClosest(points [][]int, K int) [][]int {
	q := &PointHeap{}
	result := [][]int{}

	for _, point := range points {
		heap.Push(q, point)
	}

	for i := 0; i < K; i++ {
		result = append(result, heap.Pop(q).([]int))
	}

	return result
}

// Sort
func square(x int) int {
	return x * x
}

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		d1 := square(points[i][0]) + square(points[i][1])
		d2 := square(points[j][0]) + square(points[j][1])

		return d1 < d2
	})

	return points[:K]
}
