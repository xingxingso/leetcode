package sort_colors

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
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{2, 0, 1},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "ex4",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "ex5",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: []int{0, 1, 2},
		},
	}
}

func Test_sortColors(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColors() got %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
