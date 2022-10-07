package _32_palindrome_partitioning_II

func minCut(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = -1
	for i := 1; i <= len(s); i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i-1) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)]
}

func isPalindrome(s string, i int, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
