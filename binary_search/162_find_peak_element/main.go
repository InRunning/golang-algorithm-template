package _62_find_peak_element

// 关键词：比较中点左右两边元素，往高处爬即可

// findPeakElement 左右两侧为谷底，往高处爬即可
func findPeakElement(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid-1] > nums[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] < nums[end] {
		return end
	} else {
		return start
	}
}
