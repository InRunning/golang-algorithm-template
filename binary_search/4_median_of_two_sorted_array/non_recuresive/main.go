package non_recuresive

// 关键词：两次找到第k个，取k的一半

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
	return float64(findKth(nums1, nums2, totalLen/2+1))
}

func findKth(nums1 []int, nums2 []int, k int) int {
	start1 := 0
	start2 := 0
	for {
		if start1 == len(nums1) {
			return nums2[start2+k-1]
		}
		if start2 == len(nums2) {
			return nums1[start1+k-1]
		}
		if k == 1 {
			if nums1[start1] < nums2[start2] {
				return nums1[start1]
			} else {
				return nums2[start2]
			}
		}
		half := k / 2
		mid1 := min(start1+half-1, len(nums1)-1)
		mid2 := min(start2+half-1, len(nums2)-1)
		if nums1[mid1] < nums2[mid2] {
			k -= mid1 - start1 + 1
			start1 = mid1 + 1
		} else {
			k -= mid2 - start2 + 1
			start2 = mid2 + 1
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
