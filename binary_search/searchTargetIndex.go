package binary_search

func search(nums []int, target int) int {
	// 1. 定义首位变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环，保证循环中至少有三个元素
	for start+1 < end {
		mid := (start + end) / 2
		// 3. 将中间变量与target比较
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	// 4. 检查最后剩下的两个变量
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
