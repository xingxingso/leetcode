package combinations

import (
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}},
		},
		{
			name: "ex2",
			args: args{
				n: 3,
				k: 2,
			},
			want: [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name: "ex3",
			args: args{
				n: 2,
				k: 2,
			},
			want: [][]int{{1, 2}},
		},
		{
			name: "ex4",
			args: args{
				n: 2,
				k: 1,
			},
			want: [][]int{{1}, {2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !intSliceSliceEqual(tt.want, got) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
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
