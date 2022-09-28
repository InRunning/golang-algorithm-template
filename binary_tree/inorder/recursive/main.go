package recursive

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	inOrder(root.Left)
	result = append(result, root.Val)
	inOrder(root.Right)
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
