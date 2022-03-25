package binary_search

import "testing"

func Test_search(t *testing.T) {
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
		{"中间", args{numbers, 5}, 3},
		{"左边界", args{numbers, 1}, 0},
		{"右边界", args{numbers, 10}, 8},
		{"不存在", args{numbers, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
