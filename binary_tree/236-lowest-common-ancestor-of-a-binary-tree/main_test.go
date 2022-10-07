package _36_lowest_common_ancestor_of_a_binary_tree

import (
	"reflect"
	"testing"
)

func getFirstTree() (*TreeNode, *TreeNode, *TreeNode, *TreeNode) {
	root := &TreeNode{3, nil, nil}
	left := &TreeNode{5, nil, nil}
	right := &TreeNode{1, nil, nil}
	ll := &TreeNode{6, nil, nil}
	lr := &TreeNode{2, nil, nil}
	rl := &TreeNode{0, nil, nil}
	rr := &TreeNode{8, nil, nil}
	lrl := &TreeNode{7, nil, nil}
	lrr := &TreeNode{4, nil, nil}
	root.Left = left
	root.Right = right
	root.Left.Left = ll
	root.Left.Right = lr
	root.Right.Left = rl
	root.Right.Right = rr
	root.Left.Right.Left = lrl
	root.Left.Right.Right = lrr
	return root, left, right, root
}

func getSecondTree() (*TreeNode, *TreeNode, *TreeNode, *TreeNode) {
	root := &TreeNode{3, nil, nil}
	left := &TreeNode{5, nil, nil}
	right := &TreeNode{1, nil, nil}
	ll := &TreeNode{6, nil, nil}
	lr := &TreeNode{2, nil, nil}
	rl := &TreeNode{0, nil, nil}
	rr := &TreeNode{8, nil, nil}
	lrl := &TreeNode{7, nil, nil}
	lrr := &TreeNode{4, nil, nil}
	root.Left = left
	root.Right = right
	root.Left.Left = ll
	root.Left.Right = lr
	root.Right.Left = rl
	root.Right.Right = rr
	root.Left.Right.Left = lrl
	root.Left.Right.Right = lrr
	return root, left, lrr, left
}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	firstRoot, firstP, firstQ, firstResult := getFirstTree()
	secondRoot, secondP, secondQ, secondResult := getSecondTree()
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{firstRoot, firstP, firstQ}, firstResult},

		{"2", args{secondRoot, secondP, secondQ}, secondResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
