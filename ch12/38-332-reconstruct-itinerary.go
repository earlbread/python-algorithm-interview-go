package ch12

import "sort"

func traverse(result *[]string, graph map[string][]string, dep string) {
	for len(graph[dep]) != 0 {
		nextDep := graph[dep][len(graph[dep])-1]
		graph[dep] = graph[dep][:len(graph[dep])-1]

		traverse(result, graph, nextDep)
	}

	*result = append(*result, dep)
}

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)

	// Reverse sort
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] > tickets[j][1]
		}

		return tickets[i][0] > tickets[j][0]
	})

	for _, ticket := range tickets {
		dep := ticket[0]
		arr := ticket[1]
		graph[dep] = append(graph[dep], arr)
	}

	result := make([]string, 0)

	traverse(&result, graph, "JFK")

	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return result
}
