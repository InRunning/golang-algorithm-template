package insert_sort

import (
	"reflect"
	"testing"
)

func Test_insertSort(t *testing.T) {
	type args struct {
		nums []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"反序", args{nums: []int32{5, 4, 3, 2, 1}}, []int32{1, 2, 3, 4, 5}},
		{"正序", args{nums: []int32{1, 2, 3, 4, 5}}, []int32{1, 2, 3, 4, 5}},
		{"乱序", args{nums: []int32{2, 5, 3, 4, 1}}, []int32{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
