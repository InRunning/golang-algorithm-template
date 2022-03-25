package binary_search

func search(nums []int, target int) int {
	// 1. 初始化 start 和 end
	start := 0
	end := len(nums) - 1
	// 2. 循环退出条件
	for start+1 < end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			start = middle
		} else {
			end = middle
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
