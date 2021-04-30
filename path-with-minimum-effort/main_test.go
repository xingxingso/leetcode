package path_with_minimum_effort

import "testing"

func Test_minimumEffortPath(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				heights: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			},
			want: 2,
		},
		{
			name: "equal",
			args: args{
				heights: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			},
			want: 1,
		},
		{
			name: "equal",
			args: args{
				heights: [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPath(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
