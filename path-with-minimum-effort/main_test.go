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
			name: "ex1",
			args: args{
				heights: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				heights: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				heights: [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPathO1(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPathO1() = %v, want %v", got, tt.want)
			}

			if got := minimumEffortPathO2(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPathO2() = %v, want %v", got, tt.want)
			}

			if got := minimumEffortPathO3(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPathO3() = %v, want %v", got, tt.want)
			}
		})
	}
}
