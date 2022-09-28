package _4_unique_paths

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	f := make([][]int, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if f[i] == nil {
				f[i] = make([]int, cols)
			}
			f[i][j] = 1
		}
	}
	for i := 1; i < rows; i++ {
		if obstacleGrid[i][0] == 1 || f[i-1][0] == 0 {
			f[i][0] = 0
		}
	}
	for j := 1; j < cols; j++ {
		if obstacleGrid[0][j] == 1 || f[0][j-1] == 0 {
			f[0][j] = 0
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[rows-1][cols-1]
}
