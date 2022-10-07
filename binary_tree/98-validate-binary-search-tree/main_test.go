package _8_validate_binary_search_tree

import (
	"math"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{BuildTree([]int{2, 1, 3}, 0)}, true},
		{"2", args{BuildTree([]int{5, 1, 4, math.MaxInt, math.MaxInt, 3, 6}, 0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
