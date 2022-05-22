package find_minimum_in_rotated_sorted_array_II

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
		{"中间", args{[]int{6, 6, 7, 1, 2, 2, 6, 6}}, 1},
		{"没有旋转", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"右边界", args{[]int{5, 7, 1, 1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
