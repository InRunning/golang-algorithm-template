package down_top

// 分治法：先递归再合并结果

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}
