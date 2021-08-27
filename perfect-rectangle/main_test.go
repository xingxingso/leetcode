package perfect_rectangle

import "testing"

func Test_isRectangleCover(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}},
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				rectangles: [][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}},
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4}},
			},
			want: false,
		},
		{
			name: "ex4",
			args: args{
				rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}},
			},
			want: false,
		},
		{
			name: "ex5",
			args: args{
				//[[0, 0, 1, 2], [1, 0, 3, 1], [2, 1, 3, 3], [0, 2, 2, 3], [1, 1, 2, 2]]
				rectangles: [][]int{{0, 0, 1, 2}, {1, 0, 3, 1}, {2, 1, 3, 3}, {0, 2, 2, 3}, {1, 1, 2, 2}},
			},
			want: true,
		},
		{
			name: "ex6",
			args: args{
				rectangles: [][]int{{0, 0, 3, 1}, {0, 0, 3, 1}, {0, 3, 3, 3}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleCoverP1(tt.args.rectangles); got != tt.want {
				t.Errorf("isRectangleCoverP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
