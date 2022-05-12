package _62_find_peak_element

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"最左峰值", args{nums: []int{5, 4, 3, 2, 1}}, 0},
		{"最右峰值", args{nums: []int{1, 2, 3, 4, 5}}, 4},
		{"两个峰值", args{nums: []int{1, 2, 1, 3, 5, 4}}, 4},
		{"错误", args{nums: []int{1, 2, 3, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
