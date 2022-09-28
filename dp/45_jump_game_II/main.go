package _5_jump_game_II

// 关键词：f[i] = i

func jump(nums []int) int {
	f := make([]int, len(nums))
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		f[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				f[i] = min(f[j]+1, f[i])
			}
		}
	}
	return f[len(nums)-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
