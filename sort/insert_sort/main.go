package insert_sort

func insertSort(nums []int32) []int32 {
	for index, num := range nums {
		preIndex := index - 1
		for preIndex >= 0 && nums[preIndex] > num {
			nums[preIndex+1] = nums[preIndex]
			preIndex -= 1
		}
		nums[preIndex+1] = num
	}
	return nums
}
