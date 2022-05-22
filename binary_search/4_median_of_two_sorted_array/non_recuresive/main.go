package non_recuresive

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return 0
	}
	if totalLen%2 == 0 {
		first := findKth(nums1, nums2, totalLen/2)
		second := findKth(nums1, nums2, totalLen/2+1)
		return float64(first+second) / 2
	}
	return float64(findKth(nums1, nums2, totalLen/2))
}

func findKth(nums1 []int, nums2 []int, k int) int {
	start1 := 0
	start2 := 0
	for {
		// 1. 当剩余数组长度为0时
		if start1 == len(nums1) {
			return nums2[start2+k-1]
		}
		if start2 == len(nums2) {
			return nums2[start1+k-1]
		}
		// 2. 当k为1时取两者最小值
		if k == 1 {
			if nums1[start1] < nums2[start2] {
				return nums1[start1]
			} else {
				return nums2[start2]
			}
		}
		// 3. 分治法去除一半
		half := k / 2
		index1 := min(start1+half, len(nums1)) - 1
		index2 := min(start2+half, len(nums2)) - 1
		// 3.1 比较减去小的数组长度
		if nums1[index1] < nums2[index2] {
			k -= index1 - start1 + 1
			start1 = index1 + 1
		} else {
			k -= index2 - start2 + 1
			start2 = index2 + 1
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
