package _5_jump_game

func canJump(nums []int) bool {
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if f[j] && nums[j]+j >= i {
				f[i] = true
			}
		}
	}
	return f[len(nums)-1]
}
