package search_in_rotated_array_II

func search(nums []int, target int) bool {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环条件，保证循环中至少有三个元素
	for start+1 < end {
		// 3. 比较中间变量
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		// 3.1 左边数组有序
		if nums[start] <= nums[mid] {
			if nums[mid] < target {
				start = mid
			}
		}

		// 3.2 右边数组有序
		if nums[start] > nums[mid] {
			if nums[mid] > target {

			}
		}
	}
	// 4. 处理剩下的两个变量
	if nums[start] == target {
		return true
	}
	if nums[end] == target {
		return true
	}
	return false
}
