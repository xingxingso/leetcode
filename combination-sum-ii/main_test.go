package combination_sum_ii

import (
	"sort"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name: "ex2",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: [][]int{
				{1, 2, 2},
				{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !intSliceSliceEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 二维数组是否相同， 不考虑元素顺序及子数组顺序 [[1,2]] == [[2,1]]
func intSliceSliceEqual(v1, v2 [][]int) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if len(vv1) != len(vv2) {
				continue
			}
			if sliceSameValue(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	return true
}

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
