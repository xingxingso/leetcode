package permutations

import (
	"sort"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1, 2},
			},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteO1(tt.args.nums); !intSliceSliceEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}

			if got := permuteP1(tt.args.nums); !intSliceSliceEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}

			if got := permute(tt.args.nums); !intSliceSliceEqual(tt.want, got) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
