package merge_sort

import "github.com/InRunning/golang-algorithm-template/common"

func MergeSort(nums []int) []int {
	temp := make([]int, len(nums))
	for m := 1; m < len(nums); m *= 2 {
		for i := 0; i < len(nums)-1; i += 2 * m {
			from := i
			mid := i + m - 1
			to := common.Min(i+2*m-1, len(nums)-1)
			nums = merge(nums, temp, from, mid, to)
		}
	}
	return nums
}

func merge(nums []int, temp []int, from int, mid int, to int) []int {
	i := from
	k := from
	j := mid + 1
	for i <= mid && j <= to {
		if nums[i] < nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}
	for j <= to {
		temp[k] = nums[j]
		j++
		k++
	}
	for i := from; i <= to; i++ {
		nums[i] = temp[i]
	}
	return nums
}
