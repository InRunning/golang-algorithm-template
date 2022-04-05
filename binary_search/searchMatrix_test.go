package binary_search

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"左边界", args{[][]int{{1, 2, 3}, {4, 5, 6}}, 1}, true},
		{"右边界", args{[][]int{{1, 2, 3}, {4, 5, 6}}, 6}, true},
		{"左边界不存在", args{[][]int{{1, 2, 3}, {4, 5, 6}}, -1}, false},
		{"右边界不存在", args{[][]int{{1, 2, 3}, {4, 5, 6}}, 9}, false},
		{"中间点", args{[][]int{{1, 2, 3}, {4, 5, 6}}, 3}, true},
		{"中间点不存在", args{[][]int{{1, 2, 4}, {5, 6, 7}}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
