package _53_find_minimum_in_rotated_sorted_array

// 关键词：一半连续，提前判断最左边最小值

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
		} else {
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	} else {
		return nums[end]
	}
}
