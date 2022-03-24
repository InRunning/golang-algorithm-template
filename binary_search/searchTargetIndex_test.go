package binary_search

import "testing"

func Test_searchTargetIndex(t *testing.T) {
	type args struct {
		numbers []int64
		target  int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"a", args{[]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchTargetIndex(tt.args.numbers, tt.args.target); got != tt.want {
				t.Errorf("searchTargetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
