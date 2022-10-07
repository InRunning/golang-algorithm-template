package non_recursive

import (
	"reflect"
	"testing"
)

func Test_postOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{getTree()}, []int{4, 5, 2, 6, 7, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
