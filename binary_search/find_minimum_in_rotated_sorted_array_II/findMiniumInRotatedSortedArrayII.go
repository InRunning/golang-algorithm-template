package find_minimum_in_rotated_sorted_array_II

func findMin(nums []int) int {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环条件，保证循环中至少有三个元素
	for start+1 < end {
		// 3. 比较中间变量
		mid := (start + end) / 2
		if nums[start] < nums[mid] {
			if nums[start] < nums[end] {
				return nums[start]
			} else {
				start = mid
			}
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
