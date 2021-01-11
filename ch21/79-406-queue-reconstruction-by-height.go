package ch21

import "container/heap"

type PeopleHeap [][]int

func (h PeopleHeap) Len() int { return len(h) }
func (h PeopleHeap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] > h[j][0]
}
func (h PeopleHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PeopleHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *PeopleHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func reconstructQueue(people [][]int) [][]int {
	q := &PeopleHeap{}
	for _, person := range people {
		heap.Push(q, person)
	}

	result := make([][]int, 0)
	for q.Len() > 0 {
		person := heap.Pop(q).([]int)
		index := person[1]

		result = append(result[:index], append([][]int{person}, result[index:]...)...)
	}

	return result
}
