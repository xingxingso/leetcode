package next_permutation

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "equal",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: args{
				nums: []int{1, 3, 2},
			},
		},
		{
			name: "equal",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: args{
				nums: []int{1, 2, 3},
			},
		},
		{
			name: "equal",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: args{
				nums: []int{1, 5, 1},
			},
		},
		{
			name: "equal",
			args: args{
				nums: []int{1},
			},
			want: args{
				nums: []int{1},
			},
		},
		{
			name: "equal",
			args: args{
				nums: []int{1, 3, 2},
			},
			want: args{
				nums: []int{2, 1, 3},
			},
		},
		{
			name: "equal",
			args: args{
				nums: []int{4, 5, 2, 6, 3, 1},
			},
			want: args{
				nums: []int{4, 5, 3, 1, 2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want.nums) {
				t.Errorf("%v, nextPermutation(), want %v", tt.args.nums, tt.want.nums)
			}
		})
	}
}
