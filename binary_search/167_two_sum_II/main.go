package _67_two_sum_II

// 关键词：与中点相加

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	for start+1 < end {
		mid := (start + end) / 2
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[mid] < target {
			start = mid
		} else if numbers[end]+numbers[mid] > target {
			end = mid
		} else if numbers[start]+numbers[end] < target {
			start++
		} else if numbers[start]+numbers[end] > target {
			end--
		}
	}
	if numbers[start]+numbers[end] == target {
		return []int{start + 1, end + 1}
	}
	return []int{-1, -1}
}
