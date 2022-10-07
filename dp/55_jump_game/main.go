package _5_jump_game

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
			}
		}
	}
	return dp[len(nums)-1]
}
