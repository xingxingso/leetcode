package zui_xiao_de_kge_shu_lcof

import (
	"sort"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				arr: []int{4, 5, 1, 6, 2, 7, 3, 8},
				k:   4,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "ex2",
			args: args{
				arr: []int{3, 2, 1},
				k:   2,
			},
			want: []int{1, 2},
		},
		{
			name: "ex3",
			args: args{
				arr: []int{0, 1, 2, 1},
				k:   1,
			},
			want: []int{0},
		},
		{
			name: "ex4",
			args: args{
				arr: []int{0, 1, 2, 1, 2},
				k:   3,
			},
			want: []int{0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers(tt.args.arr, tt.args.k); !sliceSameValue(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 是否元素相同
func sliceSameValue(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	t1, t2 := make([]int, len(s1)), make([]int, len(s2))
	copy(t1, s1)
	copy(t2, s2)
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	sort.Slice(t2, func(i, j int) bool {
		return t2[i] < t2[j]
	})
	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}
