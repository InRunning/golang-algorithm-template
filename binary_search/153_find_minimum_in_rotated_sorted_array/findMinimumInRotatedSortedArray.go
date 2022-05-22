package _53_find_minimum_in_rotated_sorted_array

func findMin[T int](nums []T) T {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环条件，保证循环中至少有三个元素
	for start+1 < end {
		mid := (start + end) / 2
		// 3.1 当左半段连续时
		if nums[start] < nums[mid] {
			if nums[start] < nums[end] {
				return nums[start]
			} else {
				start = mid
			}
		} else { // 3.2 当右半段连续时
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	} else {
		return nums[end]
	}
}
