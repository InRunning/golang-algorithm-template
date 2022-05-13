package _67_two_sum_II

func twoSum(numbers []int, target int) []int {
	// 1. 定义首尾变量
	start := 0
	end := len(numbers) - 1
	// 2. 保证循环中至少三个元素
	for start+1 < end {
		mid := (start + end) / 2
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[mid] == target {
			return []int{start + 1, mid + 1}
		} else if numbers[start]+numbers[mid] > target {
			end = mid
		} else if numbers[end]+numbers[mid] < target {
			start = mid
		} else if numbers[start]+numbers[end] < target {
			start++
		} else { // numbers[end] + numbers[end] > target
			end--
		}
	}
	if numbers[start]+numbers[end] == target {
		return []int{start + 1, end + 1}
	} else {
		return []int{0, 0}
	}
}
