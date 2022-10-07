package balanced_binary_tree

import "testing"

func getTree() *TreeNode {
	root := &TreeNode{3, nil, nil}
	left := &TreeNode{9, nil, nil}
	right := &TreeNode{20, nil, nil}
	rl := &TreeNode{15, nil, nil}
	rr := &TreeNode{7, nil, nil}
	root.Left = left
	root.Right = right
	root.Right.Left = rl
	root.Right.Right = rr
	return root
}

func getSecondTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	left := &TreeNode{2, nil, nil}
	right := &TreeNode{2, nil, nil}
	ll := &TreeNode{3, nil, nil}
	lr := &TreeNode{3, nil, nil}
	lll := &TreeNode{4, nil, nil}
	llr := &TreeNode{4, nil, nil}
	root.Left = left
	root.Right = right
	root.Left.Left = ll
	root.Left.Right = lr
	root.Left.Left.Left = lll
	root.Left.Left.Right = llr
	return root
}

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{getTree()}, true},
		{"2", args{getSecondTree()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
