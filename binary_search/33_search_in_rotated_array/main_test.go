package _3_search_in_rotated_array

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	numbers := []int{4, 5, 6, 7, 0, 1, 2}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{numbers, 0}, 4},
		{"3", args{numbers, 3}, -1},
		{"不存在", args{numbers, -1}, -1},
		{"左边界", args{numbers, 4}, 0},
		{"右边界", args{numbers, 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
