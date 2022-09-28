package non_recursive

// 两次循环，内部循环一直将左节点入栈，没有左节点时将右节点入栈

func inOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
