package sliding_window_maximum

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
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
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
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
				nums: []int{1, -1},
				k:    1,
			},
			want: []int{1, -1},
		},
		{
			name: "ex4",
			args: args{
				nums: []int{9, 11},
				k:    2,
			},
			want: []int{11},
		},
		{
			name: "ex5",
			args: args{
				nums: []int{4, 2},
				k:    2,
			},
			want: []int{4},
		},
		{
			name: "ex6",
			args: args{
				nums: []int{7, 2, 4},
				k:    2,
			},
			want: []int{7, 4},
		},
		{
			name: "ex7",
			args: args{
				nums: []int{1, 3, 1, 2, 0, 5},
				k:    3,
			},
			want: []int{3, 3, 2, 5},
		},
		{
			name: "ex8",
			args: args{
				nums: []int{5, 1, 3, 5, 6, 4, 8},
				k:    3,
			},
			want: []int{5, 5, 6, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}

			if got := maxSlidingWindowO1(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindowO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
