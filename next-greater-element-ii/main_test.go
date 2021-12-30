package next_greater_element_ii

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: []int{-1, 5, 5, 5, 5},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{9, 2, 6, 3, 5, 8, 1, 7},
			},
			want: []int{-1, 6, 8, 5, 8, 9, 7, 9},
		},
		{
			name: "ex4",
			args: args{
				nums: []int{2, 6, 9, 3},
			},
			want: []int{6, 9, -1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}

			if got := nextGreaterElementsP1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElementsP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
