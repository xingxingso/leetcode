package set_mismatch

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
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
				nums: []int{1, 2, 2, 4},
			},
			want: []int{2, 3},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "ex3",
			args: args{
				nums: []int{2, 3, 2},
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
