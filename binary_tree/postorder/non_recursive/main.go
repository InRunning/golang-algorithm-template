package non_recursive

// 根节点必须在右节点弹出之后，再弹出，记录上次弹出的右子节点

func getTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	left := &TreeNode{2, nil, nil}
	right := &TreeNode{3, nil, nil}
	ll := &TreeNode{4, nil, nil}
	lr := &TreeNode{5, nil, nil}
	rl := &TreeNode{6, nil, nil}
	rr := &TreeNode{7, nil, nil}
	root.Left = left
	root.Right = right
	root.Left.Left = ll
	root.Left.Right = lr
	root.Right.Left = rl
	root.Right.Right = rr
	return root
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
