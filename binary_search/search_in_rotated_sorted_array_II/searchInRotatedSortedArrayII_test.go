package search_in_rotated_sorted_array_II

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
		{"左边界", args{[]int{2, 2, 3, 3, 1, 1}, 2}, true},
		{"右边界", args{[]int{2, 2, 3, 3, 1, 1}, 1}, true},
		{"中间", args{[]int{2, 2, 3, 4, 4, 6, 0, 0, 1, 1}, 4}, true},
		{"左边界不存在", args{[]int{2, 2, 3, 3, -1, 0, 0}, 1}, false},
		{"中间不存在", args{[]int{2, 2, 3, 3, 6, 6, -1, -1, 0, 0}, 5}, false},
		{"上一次的错误答案", args{[]int{1, 0, 1, 1, 1}, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
