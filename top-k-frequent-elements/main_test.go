package top_k_frequent_elements

import (
	"sort"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{1, 1, 1, 1, 2, 2, 3, 4},
				k:    2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !sliceSameValue(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}

			if got := topKFrequentS2(tt.args.nums, tt.args.k); !sliceSameValue(got, tt.want) {
				t.Errorf("topKFrequentS2() = %v, want %v", got, tt.want)
			}

			if got := topKFrequentO1(tt.args.nums, tt.args.k); !sliceSameValue(got, tt.want) {
				t.Errorf("topKFrequentO1() = %v, want %v", got, tt.want)
			}
		})
	}
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
