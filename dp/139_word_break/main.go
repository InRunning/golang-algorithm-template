package _39_word_break

// 参考：https://www.cnblogs.com/du001011/p/11259100.html

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && inDict(s[j:i], wordDict) {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func inDict(s string, wordDict []string) bool {
	for i := 0; i < len(wordDict); i++ {
		if s == wordDict[i] {
			return true
		}
	}
	return false
}
