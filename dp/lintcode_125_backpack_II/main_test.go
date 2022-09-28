package lintcode_125_backpack_II

import "testing"

func TestBackPackII(t *testing.T) {
	type args struct {
		m int
		a []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10, []int{2, 3, 5, 7}, []int{1, 5, 2, 4}}, 9},
		{"2", args{10, []int{2, 3, 8}, []int{2, 5, 8}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackPackII(tt.args.m, tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("BackPackII() = %v, want %v", got, tt.want)
			}
		})
	}
}
