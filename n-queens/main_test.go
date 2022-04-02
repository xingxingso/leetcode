package n_queens

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "ex1",
			args: args{
				n: 4,
			},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name: "ex2",
			args: args{
				n: 1,
			},
			want: [][]string{{"Q"}},
		},
		{
			name: "ex3",
			args: args{
				n: 6,
			},
			want: [][]string{
				{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."},
				{"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."},
				{"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."},
				{"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
			if got := solveNQueensO1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
			if got := solveNQueensO2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}

			if got := solveNQueensP1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
