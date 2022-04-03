package binary_search

import "sync"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	var leftIndex int
	var rightIndex int
	go func() {
		defer waitGroup.Done()
		leftIndex = binarySearch(nums, target, true)
	}()
	go func() {
		defer waitGroup.Done()
		rightIndex = binarySearch(nums, target, false)
	}()
	waitGroup.Wait()
	if leftIndex == -1 {
		return []int{-1, -1}
	}
	return []int{leftIndex, rightIndex}
}

func binarySearch(nums []int, target int, lower bool) int {
	// 1. 定义首尾变量
	start := 0
	end := len(nums) - 1
	// 2. 定义循环条件，保证循环中至少有三个元素
	for start+1 < end {
		// 3. 比较中间变量
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		} else {
			if lower {
				end = mid
			} else {
				start = mid
			}
		}
	}
	// 4. 处理首尾变量
	if lower {
		if nums[start] == target {
			return start
		} else if nums[end] == target {
			return end
		}
		return -1
	} else {
		if nums[end] == target {
			return end
		} else if nums[start] == target {
			return start
		}
		return -1
	}
}
