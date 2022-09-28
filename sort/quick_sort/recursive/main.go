package recursive

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}

func quicksort(nums []int, start int, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quicksort(nums, start, pivot-1)
		quicksort(nums, pivot+1, end)
	}
}

func partition(nums []int, start int, end int) int {
	// 最后一个元素作为基准
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
