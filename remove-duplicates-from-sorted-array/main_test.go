package remove_duplicates_from_sorted_array

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
				nums: []int{1, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{},
			},
			want: []int{},
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
				nums: []int{1, 1},
			},
			want: []int{1},
		},
		{
			name: "ex6",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
}

func Test_removeDuplicates(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != len(tt.want) || !reflect.DeepEqual(tt.args.nums[:got], tt.want) {
				t.Errorf("removeDuplicates(), nums=%v, want %v", tt.args.nums[:got], tt.want)
			}
		})
	}
}

func Test_removeDuplicatesS1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesS1(tt.args.nums); got != len(tt.want) || !reflect.DeepEqual(tt.args.nums[:got], tt.want) {
				t.Errorf("removeDuplicatesS1(), nums=%v, want %v", tt.args.nums[:got], tt.want)
			}
		})
	}
}
