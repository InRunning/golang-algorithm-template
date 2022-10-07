package top_down

import (
	"reflect"
	"testing"
)

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

func Test_preOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{getTree()}, []int{1, 2, 4, 5, 3, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
