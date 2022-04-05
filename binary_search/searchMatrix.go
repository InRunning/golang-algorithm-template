package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	// 1. 定义首尾变量
	rows := len(matrix)
	columns := len(matrix[0])
	start := 0
	end := rows*columns - 1
	// 2. 定义循环条件，保证循环中至少有三个元素
	for start+1 < end {
		// 3. 比较中间变量
		mid := (start + end) / 2
		if matrix[mid/columns][mid%columns] < target {
			start = mid
		} else if matrix[mid/columns][mid%columns] > target {
			end = mid
		} else {
			return true
		}
	}
	// 处理剩下的两个变量
	if matrix[start/columns][start%columns] == target || matrix[end/columns][end%columns] == target {
		return true
	}
	return false
}
