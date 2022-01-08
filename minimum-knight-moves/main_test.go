package minimum_knight_moves

import "testing"

func Test_minKnightMoves(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				x: 2,
				y: 1,
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				x: 5,
				y: 5,
			},
			want: 4,
		},
		{
			name: "ex3",
			args: args{
				x: 6,
				y: 5,
			},
			want: 5,
		},
		{
			name: "ex4",
			args: args{
				x: 30,
				y: 15,
			},
			want: 15,
		},
		{
			name: "ex5",
			args: args{
				x: 87,
				y: 79,
			},
			want: 56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKnightMovesP1(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("minKnightMoves() = %v, want %v", got, tt.want)
			}
			if got := minKnightMovesP2(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("minKnightMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
