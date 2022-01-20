package the_skyline_problem

import (
	"reflect"
	"testing"
)

func Test_getSkyline(t *testing.T) {
	type args struct {
		buildings [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				buildings: [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			},
			want: [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{
			name: "ex2",
			args: args{
				buildings: [][]int{{2, 9, 10}, {2, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			},
			want: [][]int{{2, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{
			name: "ex3",
			args: args{
				buildings: [][]int{{2, 9, 10}, {2, 7, 15}, {5, 7, 12}, {15, 20, 10}, {19, 24, 8}},
			},
			want: [][]int{{2, 15}, {7, 10}, {9, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{
			name: "ex4",
			args: args{
				buildings: [][]int{{0, 2, 3}, {2, 5, 3}},
			},
			want: [][]int{{0, 3}, {5, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSkyline(tt.args.buildings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}
