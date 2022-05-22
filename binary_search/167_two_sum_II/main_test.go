package _67_two_sum_II

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"三个", args{[]int{2, 3, 4}, 6}, []int{1, 3}},
		{"两个", args{[]int{-1, 0}, -1}, []int{1, 2}},
		{"5个", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 15}, []int{7, 8}},
		{"5个", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
