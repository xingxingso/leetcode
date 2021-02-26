package transpose_matrix

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		/*{
			name: "equal0",
			args: args{
				matrix: [][]int{},
			},
			want: nil,
		},*/
		{
			name: "equal1",
			args: args{
				matrix: [][]int{{1}},
			},
			want: [][]int{{1}},
		},
		{
			name: "equal2",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "equal3",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}

			if got := transposeO1(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
