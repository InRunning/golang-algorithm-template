package binary_search

func mySqrt(x int) int {
	// 1. 定义首尾变量
	start := 1
	end := x
	// 2. 定义循环条件，保证循环中至少三个元素
	for start+1 < end {
		// 3. 比较中间变量
		mid := (start + end) / 2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			start = mid
		} else {
			end = mid
		}
	}
	if end > x/end {
		return start
	}
	return end
}
