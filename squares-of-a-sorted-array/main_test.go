package squares_of_a_sorted_array

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
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
	}
}

func Test_sortedSquares(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquaresO1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquaresO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
