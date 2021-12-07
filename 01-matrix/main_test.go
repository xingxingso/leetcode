package _01_matrix

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "ex2",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
