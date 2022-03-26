package binary_search

func searchInsert(nums []int, target int) int {
	// 1. 定义首位变量
	start := 0
	end := len(nums) - 1
	// 2. 循环条件，保证循环中至少三个元素
	for start+1 < end {
		// 3. 定义中间变量并进行比较
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] >= target {
		return start
	}
	if nums[end] >= target {
		return end
	}
	return end + 1
}
