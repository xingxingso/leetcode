package rotate_image

import (
	"reflect"
	"testing"
)

type args struct {
	matrix [][]int
}

func getTests() []struct {
	name string
	args args
	want [][]int
} {
	return []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "ex2",
			args: args{
				matrix: [][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "ex3",
			args: args{
				matrix: [][]int{
					{1},
				},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "ex4",
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},
	}
}

func Test_rotate(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if rotate(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("rotate() got %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
