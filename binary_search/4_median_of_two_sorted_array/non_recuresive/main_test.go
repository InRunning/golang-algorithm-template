package non_recuresive

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"奇数", args{[]int{1, 2, 3}, []int{2, 5}}, 2},
		{"偶数", args{[]int{1, 3, 5}, []int{2, 4, 6}}, 3.5},
		{"3", args{[]int{1, 3}, []int{2}}, 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
