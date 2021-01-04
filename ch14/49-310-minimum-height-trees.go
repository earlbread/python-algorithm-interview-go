package ch14

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := map[int][]int{}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	leaves := []int{}
	for k, v := range graph {
		if len(v) == 1 {
			leaves = append(leaves, k)
		}
	}

	for len(leaves) < n {
		n -= len(leaves)

		new_leaves := []int{}

		for _, leaf := range leaves {
			node := graph[leaf][0]

			// Remove leaf from node
			for i := 0; i < len(graph[node]); i++ {
				if graph[node][i] == leaf {
					graph[node][i] = graph[node][len(graph[node])-1]
					break
				}
			}
			graph[node] = graph[node][:len(graph[node])-1]

			if len(graph[node]) == 1 {
				new_leaves = append(new_leaves, node)
			}
		}

		leaves = new_leaves
	}

	return leaves
}
