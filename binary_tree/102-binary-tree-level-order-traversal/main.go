package _02_binary_tree_level_order_traversal

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			root = queue[0]
			queue = queue[1:]
			temp = append(temp, root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		result = append(result, temp)
	}
	return result
}
