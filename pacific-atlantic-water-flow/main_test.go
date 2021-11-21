package pacific_atlantic_water_flow

import (
	"reflect"
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex0",
			args: args{
				heights: [][]int{
					{1, 2, 2, 3, 5},
					{3, 2, 3, 4, 4},
					{2, 4, 5, 3, 1},
					{6, 7, 1, 4, 5},
					{5, 1, 1, 2, 4},
				},
			},
			want: [][]int{
				{0, 4},         //5
				{1, 3}, {1, 4}, //4,4
				{2, 2},         //5
				{3, 0}, {3, 1}, //6,7
				{4, 0}, //5
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pacificAtlantic(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pacificAtlantic() = %v, want %v", got, tt.want)
			}

			if got := pacificAtlanticP1(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pacificAtlanticP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
