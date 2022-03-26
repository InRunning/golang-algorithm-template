package binary_search

func searchInRotatedSortedArray(nums []int, target int) int {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		// 2. 定义中间变量并分四种情况进行比较
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		// 2.1 左半段是连续的，注意：左半段和右半段必然有一段是连续的
		if nums[start] < nums[mid] {
			// 2.1.1 target在左半段中间
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else { // 2.2 右半段是连续
			// 2.2.1 target 在右半段中间
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
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
