package merge_sorted_array

import (
	"reflect"
	"testing"
)

type args struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func getTests() []struct {
	name string
	args args
	want []int
} {
	return []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "ex2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "ex3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
	}
}

func Test_merge(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("merge() got %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}

func Test_mergeS1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if mergeS2(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("mergeS2() got %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
