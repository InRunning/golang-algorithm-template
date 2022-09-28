package lintcode_92_backpack

import "testing"

func TestBackPack(t *testing.T) {
	type args struct {
		m int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10, []int{3, 4, 8, 5}}, 9},
		{"2", args{12, []int{2, 3, 5, 7}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackPack(tt.args.m, tt.args.a); got != tt.want {
				t.Errorf("BackPack() = %v, want %v", got, tt.want)
			}
		})
	}
}
