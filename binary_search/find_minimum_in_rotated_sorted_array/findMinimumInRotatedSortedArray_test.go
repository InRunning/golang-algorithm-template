package find_minimum_in_rotated_sorted_array

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
		{"右边中间", args{[]int{3, 4, 5, 6, 7, 1, 2}}, 1},
		{"左边中间", args{[]int{5, 6, 1, 2, 3, 4}}, 1},
		{"没有旋转", args{[]int{1, 2, 3, 4, 5, 6}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
