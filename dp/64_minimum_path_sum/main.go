package _4_minimum_path_sum

func minPathSum(grid [][]int) int {
	if len(grid) == 0 && len(grid[0]) == 0 {
		return 0
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
