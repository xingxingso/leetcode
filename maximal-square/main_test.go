package maximal_square

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			want: 4,
		},
		{
			name: "ex2",
			args: args{
				matrix: [][]byte{
					{'0', '1'},
					{'1', '0'},
				},
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				matrix: [][]byte{
					{'0'},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquareP1(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquareP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
