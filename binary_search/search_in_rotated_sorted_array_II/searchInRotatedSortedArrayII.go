package search_in_rotated_sorted_array_II

// search 该题由于可能全部元素都重复，所以最终会退化成O(n)时间复杂度
func search(nums []int, target int) bool {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环条件，保证循环中至少有三个元素
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		// 3. 左半段连续
		if nums[start] < nums[mid] {
			// 3.1 如果在左半段中间
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else { //
				start = mid
			}
		} else if nums[start] > nums[mid] { // 左半段不连续
			// 3.2 不在左半段中间，即在旋转数组中间
			if nums[mid] <= target && target < nums[start] {
				start = mid
			} else {
				end = mid
			}
		} else { // 相等时只能遍历
			start++
		}
	}
	if nums[start] == target {
		return true
	}
	if nums[end] == target {
		return true
	}
	return false
}
