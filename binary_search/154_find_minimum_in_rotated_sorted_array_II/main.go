package _54_find_minimum_in_rotated_sorted_array_II

// 关键词：相等时遍历

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[start] < nums[end] {
			return nums[start]
		}
		if nums[start] < nums[mid] {
			start = mid
		} else if nums[mid] < nums[end] {
			end = mid
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
