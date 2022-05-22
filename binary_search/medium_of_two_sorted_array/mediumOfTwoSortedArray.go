package medium_of_two_sorted_array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	// 奇数时取中间一个数
	if totalLength%2 == 1 {
		return float64(findKthElement(nums1, nums2, totalLength/2+1))
	} else {
		leftMediumIndex := totalLength/2 - 1
		rightMediumIndex := totalLength / 2
		mediumSum := findKthElement(nums1, nums2, leftMediumIndex+1) + findKthElement(nums1, nums2, rightMediumIndex+1)
		return float64(mediumSum) / 2.0
	}
}

func findKthElement(nums1 []int, nums2 []int, k int) int {
	startIndex1, startIndex2 := 0, 0
	for {
		// 1. 处理特殊情况
		// 1.1 数组长度为0时
		if startIndex1 == len(nums1) {
			return nums2[startIndex2+k-1]
		}
		if startIndex2 == len(nums2) {
			return nums1[startIndex1+k-1]
		}
		// 1.2 当k为1时取两个首位最小数
		if k == 1 {
			return min(nums1[startIndex1], nums2[startIndex2])
		}
		// 1.3 k越界时
		half := k / 2
		index1 := min(startIndex1+half, len(nums1)) - 1
		index2 := min(startIndex2+half, len(nums2)) - 1
		// 2. 去除小的数组左边
		if nums1[index1] <= nums2[index2] {
			k -= index1 - startIndex1 + 1
			startIndex1 = index1 + 1
		} else {
			k -= index2 - startIndex2 + 1
			startIndex2 = index2 + 1
		}
	}
}

func min(first int, second int) int {
	if first <= second {
		return first
	} else {
		return second
	}
}
