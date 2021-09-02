package next_permutation

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

func getTests() []struct {
	name string
	args args
	want args
} {
	return []struct {
		name string
		args args
		want args
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: args{
				nums: []int{1, 3, 2},
			},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: args{
				nums: []int{1, 2, 3},
			},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: args{
				nums: []int{1, 5, 1},
			},
		},
		{
			name: "ex4",
			args: args{
				nums: []int{1},
			},
			want: args{
				nums: []int{1},
			},
		},
		{
			name: "ex5",
			args: args{
				nums: []int{1, 3, 2},
			},
			want: args{
				nums: []int{2, 1, 3},
			},
		},
		{
			name: "ex6",
			args: args{
				nums: []int{4, 5, 2, 6, 3, 1},
			},
			want: args{
				nums: []int{4, 5, 3, 1, 2, 6},
			},
		},
		{
			name: "ex7",
			args: args{
				nums: []int{1, 5, 1},
			},
			want: args{
				nums: []int{5, 1, 1},
			},
		},
	}
}

func Test_nextPermutation(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want.nums) {
				t.Errorf("%v, nextPermutation(), want %v", tt.args.nums, tt.want.nums)
			}
		})
	}
}

func Test_nextPermutationO1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutationO1(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want.nums) {
				t.Errorf("%v, nextPermutationO1(), want %v", tt.args.nums, tt.want.nums)
			}
		})
	}
}
