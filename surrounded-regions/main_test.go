package surrounded_regions

import (
	"reflect"
	"testing"
)

type args struct {
	board [][]byte
}

func getTests() []struct {
	name string
	args args
	want [][]byte
} {
	return []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "ex1",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "ex2",
			args: args{
				board: [][]byte{{'X'}},
			},
			want: [][]byte{{'X'}},
		},
	}
}

func Test_solve(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if solve(tt.args.board); !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}

func Test_solveO1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if solveO1(tt.args.board); !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("solveO1() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}

func Test_solveO2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if solveO2(tt.args.board); !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("solveO2() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
