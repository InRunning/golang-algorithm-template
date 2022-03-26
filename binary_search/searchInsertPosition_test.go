package binary_search

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	numbers := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"中间", args{numbers, 6}, 4},
		{"左边界", args{numbers, 1}, 0},
		{"右边界", args{numbers, 10}, 8},
		{"不存在", args{numbers, 3}, 2},
		{"右边界不存在", args{numbers, 11}, 9},
		{"左边界不存在", args{numbers, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
