package factor_combinations

import (
	"testing"
)

func Test_getFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "equal",
			args: args{
				n: 1,
			},
			want: [][]int{},
		},
		{
			name: "equal",
			args: args{
				n: 37,
			},
			want: [][]int{},
		},
		{
			name: "equal",
			args: args{
				n: 12,
			},
			want: [][]int{{2, 6}, {2, 2, 3}, {3, 4}},
		},
		{
			name: "equal",
			args: args{
				n: 32,
			},
			want: [][]int{{2, 16}, {2, 2, 8}, {2, 2, 2, 4}, {2, 2, 2, 2, 2}, {2, 4, 4}, {4, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFactors(tt.args.n); !intSliceSliceEqual(got, tt.want) {
				t.Errorf("getFactors() = %v, want %v", got, tt.want)
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
loop:
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				continue loop
			}
		}
		return false
	}

	return true
}
