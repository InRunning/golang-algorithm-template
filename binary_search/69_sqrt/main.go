package _9_sqrt

// 关键词：先判断尾部

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	start := 1
	end := x
	for start+1 < end {
		mid := (start + end) / 2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			start = mid
		} else {
			end = mid
		}
	}
	if end == x/end {
		return end
	} else {
		return start
	}
}
