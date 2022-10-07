package maximum_depth_of_binary_tree

import "testing"

func getTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	left := &TreeNode{2, nil, nil}
	right := &TreeNode{3, nil, nil}
	rr := &TreeNode{7, nil, nil}
	root.Left = left
	root.Right = right
	root.Right.Right = rr
	return root
}

func getSecondTree() *TreeNode {
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

func getThirdTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	right := &TreeNode{2, nil, nil}
	root.Right = right
	return root
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{getTree()}, 3},
		{"2", args{getSecondTree()}, 3},
		{"3", args{getThirdTree()}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
