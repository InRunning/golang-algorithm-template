package lintcode_125_backpack_II

func BackPackII(m int, a []int, v []int) int {
	dp := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0
	for i := 1; i <= len(a); i++ {
		for j := 0; j <= m; j++ {
			if j >= a[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-a[i-1]]+v[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(a)][m]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
