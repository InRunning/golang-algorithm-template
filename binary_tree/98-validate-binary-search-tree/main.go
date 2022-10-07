package _8_validate_binary_search_tree

// 二叉搜索树中序遍历

func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}
