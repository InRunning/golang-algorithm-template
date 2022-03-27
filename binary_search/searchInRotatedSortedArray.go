package binary_search

func searchInRotatedSortedArray(nums []int, target int) int {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环退出条件，保证循环中至少有三个元素
	for start+1 < end {
		// 3. 将中间变量与target比较，注意左半段和右半段必然有一段连续
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		// 3.1 左半段连续时
		if nums[start] <= nums[mid] {
			// 3.1.1 变量在左半段中间
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		}
		// 3.2 右半段连续时
		if nums[mid] <= nums[end] {
			// 3.2.1 变量在右半段中间
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	// 4. 比较剩下的首尾变量
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
