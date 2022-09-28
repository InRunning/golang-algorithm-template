package _32_palindrome_partitioning_II

func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}
	f := make([]int, len(s)+1)
	f[0] = -1
	f[1] = 0
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i-1) {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[len(s)]
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
