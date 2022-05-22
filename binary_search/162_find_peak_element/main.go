package _62_find_peak_element

// findPeakElement 左右两侧为谷底，往高处爬即可
func findPeakElement(nums []int) int {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 保证循环中至少有三个元素
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] < nums[mid+1] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return end
	} else {
		return start
	}
}
