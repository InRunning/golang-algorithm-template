package search_in_rotated_array_II

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"左边界", args{[]int{3, 3, 4, 4, 1, 1, 2}, 3}, true},
		{"右边界", args{[]int{3, 3, 4, 4, 1, 1, 2}, 2}, true},
		{"左边界不存在", args{[]int{4, 4, 5, 6, 1, 2, 2}, 3}, false},
		{"中间点", args{[]int{3, 3, 4, 5, 1, 1, 2}, 4}, true},
		{"中间点不存在", args{[]int{3, 3, 4, 6, 7, 1, 2, 2, 3}, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
