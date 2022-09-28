package lintcode_92_backpack

// 题解：https://zhengyang2015.gitbooks.io/lintcode/content/backpack_i_92.html

// 不要第i件物品与要第i件物品
// dp[i][j] 表示第i件物品是否刚好装满j的容量

func BackPack(m int, a []int) int {
	dp := make([][]bool, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(a); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j-a[i-1] >= 0 && dp[i-1][j-a[i-1]] {
				dp[i][j] = true
			}
		}
	}
	for i := m; i >= 0; i-- {
		if dp[len(a)][i] {
			return i
		}
	}
	return 0
}
