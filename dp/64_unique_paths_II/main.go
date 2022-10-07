package _4_unique_paths_II

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	dp[0][0] = 1
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] != 1 && dp[i-1][0] != 0 {
			dp[i][0] = 1
		}
	}
	for j := 1; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] != 1 && dp[0][j-1] != 0 {
			dp[0][j] = 1
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
