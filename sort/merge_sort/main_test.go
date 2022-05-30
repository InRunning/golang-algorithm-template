package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"反序", args{nums: []int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"正序", args{nums: []int{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
		{"乱序", args{nums: []int{2, 5, 3, 4, 1}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		array []int
		temp  []int
		from  int
		mid   int
		to    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"测试", args{[]int{2, 1, 4, 3, 5}, []int{0, 0, 0, 0, 0}, 0, 0, 1}, []int{1, 2, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.array, tt.args.temp, tt.args.from, tt.args.mid, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
