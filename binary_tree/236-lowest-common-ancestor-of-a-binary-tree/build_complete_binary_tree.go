package _36_lowest_common_ancestor_of_a_binary_tree

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
