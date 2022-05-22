package _54_find_minimum_in_rotated_sorted_array_II

func findMin(nums []int) int {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环条件，保证循环中至少有三个元素
	for start+1 < end {
		// 3.1 如果剩余数组为升序
		if nums[start] < nums[end] {
			return nums[start]
		}
		mid := (start + end) / 2
		// 3.2 如果最小值在左边，即左半边不连续
		if nums[start] > nums[mid] {
			end = mid
		} else if nums[mid] > nums[end] { // 最小值在右边
			start = mid + 1
		} else {
			start++
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	} else {
		return nums[end]
	}
}
