package remove_element

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
	val  int
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
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: []int{2, 2},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: []int{0, 1, 3, 0, 4},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{0},
				val:  2,
			},
			want: []int{0},
		},
		{
			name: "ex4",
			args: args{
				nums: []int{0},
				val:  0,
			},
			want: []int{},
		},
		{
			name: "ex5",
			args: args{
				nums: []int{},
				val:  0,
			},
			want: []int{},
		},
	}
}

func Test_removeElement(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != len(tt.want) || !reflect.DeepEqual(tt.args.nums[:got], tt.want) {
				t.Errorf("removeDuplicates(), nums=%v, want %v", tt.args.nums[:got], tt.want)
			}
		})
	}
}
