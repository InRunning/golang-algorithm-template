package merge_sort

import "github.com/InRunning/golang-algorithm-template/common"

func MergeSort(nums []int) []int {
	temp := make([]int, len(nums))
	// m = 1时，i = 0, 2, 4
	for m := 1; m < len(nums); m *= 2 {
		// i < len(nums) -1 可以避免只排序一个元素（最后一个）
		for i := 0; i < len(nums)-1; i += 2 * m {
			from := i
			// 中间是一个虚拟分界，使用左半边最末端的索引代替
			mid := i + m - 1
			to := common.Min(i+2*m-1, len(nums)-1)
			nums = merge(nums, temp, from, mid, to)
		}
	}
	return nums
}

func merge(array []int, temp []int, from int, mid int, to int) []int {
	k := from
	i := from
	j := mid + 1
	for i <= mid && j <= to {
		if array[i] < array[j] {
			temp[k] = array[i]
			i++
		} else {
			temp[k] = array[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = array[i]
		k++
		i++
	}
	for j <= to {
		temp[k] = array[j]
		k++
		j++
	}
	for i = from; i <= to; i++ {
		array[i] = temp[i]
	}
	return array
}
