package _3_search_in_rotated_array

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[start] < nums[mid] { // 左半边数组连续
			if nums[start] <= target && nums[mid] >= target { // 目标在左边数组
				end = mid
			} else {
				start = mid
			}
		} else { // 右半边数组连续
			if nums[mid] <= target && nums[end] >= target {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}
