package sliding_puzzle

import "testing"

func Test_slidingPuzzle(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				board: [][]int{{1, 2, 3}, {4, 0, 5}},
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				board: [][]int{{1, 2, 3}, {5, 4, 0}},
			},
			want: -1,
		},
		{
			name: "ex3",
			args: args{
				board: [][]int{{4, 1, 2}, {5, 0, 3}},
			},
			want: 5,
		},
		{
			name: "ex4",
			args: args{
				board: [][]int{{3, 2, 4}, {1, 5, 0}},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingPuzzle(tt.args.board); got != tt.want {
				t.Errorf("slidingPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
