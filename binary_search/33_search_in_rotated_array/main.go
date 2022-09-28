package _3_search_in_rotated_array

// 关键词：有一边连续

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[start] < nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else {
			if nums[mid] <= target && target <= nums[end] {
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
	} else {
		return -1
	}
}
