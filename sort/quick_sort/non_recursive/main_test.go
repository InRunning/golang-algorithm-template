package non_recursive

import (
	"reflect"
	"testing"
)

func Test_quickSortIterative(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"正序", args{nums: []int{1, 2, 3, 4, 5, 6, 7}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"反序", args{nums: []int{7, 6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"乱序", args{nums: []int{3, 4, 2, 1, 7, 6, 5}}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortIterative(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
