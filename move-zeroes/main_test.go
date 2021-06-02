package move_zeroes

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
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{},
			},
			want: []int{},
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
	}
}

func Test_moveZeroes(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes(), nums=%v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_moveZeroesO1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroesO1(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroesO1(), nums=%v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
