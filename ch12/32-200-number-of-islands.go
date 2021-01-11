package ch12

func removeIsland(grid [][]byte, row, col int) {
	rows := len(grid)
	cols := len(grid[0])

	if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] != '1' {
		return
	}

	grid[row][col] = '0'
	removeIsland(grid, row+1, col)
	removeIsland(grid, row-1, col)
	removeIsland(grid, row, col+1)
	removeIsland(grid, row, col-1)
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				removeIsland(grid, i, j)
			}
		}
	}

	return count
}
