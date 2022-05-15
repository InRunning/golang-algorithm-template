package _4_find_first_and_last_position_of_element_in_sorted_array

import "sync"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var start int
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		start = search(nums, target, true)
	}()
	end := search(nums, target, false)
	waitGroup.Wait()
	return []int{start, end}
}

// search 两次二分，isFirst: 是否为首位置
func search(nums []int, target int, isFirst bool) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if isFirst {
				end = mid
			} else {
				start = mid
			}
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	if isFirst {
		if nums[start] == target {
			return start
		} else if nums[end] == target {
			return end
		} else {
			return -1
		}
	} else {
		if nums[end] == target {
			return end
		} else if nums[start] == target {
			return start
		} else {
			return -1
		}
	}
}
