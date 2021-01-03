package ch13

import "container/heap"

type RouteHeap []Route

func (h RouteHeap) Len() int           { return len(h) }
func (h RouteHeap) Less(i, j int) bool { return h[i].Price < h[j].Price }
func (h RouteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RouteHeap) Push(x interface{}) {
	*h = append(*h, x.(Route))
}

func (h *RouteHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

type Price struct {
	V int
	W int
}

type Route struct {
	Price   int
	Stop    int
	Arrival int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := map[int][]Price{}

	for _, flight := range flights {
		u, v, w := flight[0], flight[1], flight[2]
		graph[u] = append(graph[u], Price{v, w})
	}

	q := &RouteHeap{Route{0, -1, src}}

	for q.Len() > 0 {
		route := heap.Pop(q).(Route)

		if route.Stop > K {
			continue
		}

		if route.Arrival == dst {
			return route.Price
		}

		for _, p := range graph[route.Arrival] {
			heap.Push(q, Route{p.W + route.Price, route.Stop + 1, p.V})
		}
	}

	return -1
}
