package binary_search

type Number interface {
	int64 | float64
}

func searchTargetIndex[N Number](numbers []N, target N) N {
	//1. 定义首尾变量
	start := 0
	end := len(numbers) - 1
	//2. 定义循环条件，保证循环时至少有三个元素
	for start+1 < end {
		//3. 定义中间变量并进行比较
		mid := (start + end) / 2
		if numbers[mid] == target {
			return N(mid)
		} else if numbers[mid] < target {
			mid = start
		} else {
			mid = end
		}
	}
	//4. 判断最后一个元素是否等于target
	if numbers[start] == target {
		return N(start)
	}
	if numbers[end] == target {
		return N(end)
	}
	return N(-1)
}
