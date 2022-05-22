package quick_sort

import stack2 "github.com/InRunning/golang-algorithm-template/stack"

func quickSortIterative(nums []int) []int {
	stack := stack2.NewStack()
	// 将首尾索引放入栈中
	stack.Push(0)
	stack.Push(len(nums) - 1)
	for stack.Top != nil {
		endIndex := stack.Pop().(int)
		startIndex := stack.Pop().(int)
		pivotIndex := partition(nums, startIndex, endIndex)
		// 如果左边存在两个以上元素
		if (pivotIndex - 1) > startIndex {
			stack.Push(startIndex)
			stack.Push(pivotIndex - 1)
		}
		if (pivotIndex + 1) < endIndex {
			stack.Push(pivotIndex + 1)
			stack.Push(endIndex)
		}
	}
	return nums
}

// 将小于pivot的元素置于pivot左边，将大于pivot的元素置于pivot右边
func partition(nums []int, startIndex int, endIndex int) int {
	// 选取最后一个元素作为pivot
	pivot := nums[endIndex]
	for j := startIndex; j <= endIndex-1; j++ {
		if nums[j] <= pivot {
			temp := nums[startIndex]
			nums[startIndex] = nums[j]
			nums[j] = temp
			startIndex++
		}
	}
	temp := nums[startIndex]
	nums[startIndex] = pivot
	nums[endIndex] = temp
	return startIndex
}
