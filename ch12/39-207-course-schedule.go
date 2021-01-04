package ch12

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

	traced := map[int]bool{}
	visited := map[int]bool{}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i, graph, traced, visited) {
			return false
		}
	}

	return true
}

func hasCycle(course int, graph map[int][]int, traced, visited map[int]bool) bool {
	if traced[course] {
		return true
	}

	if visited[course] {
		return false
	}

	traced[course] = true

	for _, pre := range graph[course] {
		if hasCycle(pre, graph, traced, visited) {
			return true
		}
	}

	traced[course] = false
	visited[course] = true

	return false
}
