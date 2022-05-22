package _54_find_minimum_in_rotated_sorted_array_II

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"左半边最小值", args{[]int{5, 6, 1, 2, 3, 4, 4}}, 1},
		{"右半边最小值", args{[]int{3, 4, 5, 6, 1, 2, 3}}, 1},
		{"重复值", args{[]int{1, 1, 1, 1, 1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
