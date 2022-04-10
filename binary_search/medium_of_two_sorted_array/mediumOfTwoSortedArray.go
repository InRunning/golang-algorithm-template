package medium_of_two_sorted_array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	// 奇数时只需取中间的数
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(findKthElement(nums1, nums2, midIndex+1))
	} else { // 偶数时取两个数
		midIndex1 := totalLength/2 - 1
		midIndex2 := totalLength / 2
		mediumNumSum := findKthElement(nums1, nums2, midIndex1+1) + findKthElement(nums1, nums2, midIndex2+1)
		return float64(mediumNumSum) / 2.0
	}
}

func findKthElement(nums1 []int, nums2 []int, k int) int {
	startIndex1, startIndex2 := 0, 0
	for {
		// 1. 处理特殊情况
		// 1.1 当数组长度为0时
		if startIndex1 == len(nums1) {
			return nums2[startIndex2+k-1]
		}
		if startIndex2 == len(nums2) {
			return nums1[startIndex1+k-1]
		}
		// 1.2 当k为1时取两者首尾最小值
		if k == 1 {
			return min(nums1[startIndex1], nums2[startIndex2])
		}
		// 1.3 当k越界时
		half := k / 2
		index1 := min(startIndex1+half, len(nums1)) - 1
		index2 := min(startIndex2+half, len(nums2)) - 1
		// 2. 比较后减去小的数组长度
		if nums1[index1] < nums2[index2] {
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
