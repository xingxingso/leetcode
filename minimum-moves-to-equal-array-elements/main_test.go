package minimum_moves_to_equal_array_elements

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 0,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{7, 5, 1, 9, 10},
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.nums); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}

			if got := minMovesS2(tt.args.nums); got != tt.want {
				t.Errorf("minMovesS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
