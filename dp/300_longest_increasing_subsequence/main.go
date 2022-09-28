package _00_longest_increasing_subsequence

// 参考解释：https://www.cnblogs.com/du001011/p/11259100.html

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	result := f[0]
	for i := 1; i < len(nums); i++ {
		result = max(result, f[i])
	}
	return result
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
