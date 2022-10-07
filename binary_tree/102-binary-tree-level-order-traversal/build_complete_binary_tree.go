package _02_binary_tree_level_order_traversal

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(array []int, i int) *TreeNode {
	if i < len(array) {
		if array[i] == math.MaxInt {
			return nil
		}
		root := &TreeNode{array[i], nil, nil}
		root.Left = BuildTree(array, 2*i+1)
		root.Right = BuildTree(array, 2*i+2)
		return root
	} else {
		return nil
	}
}
