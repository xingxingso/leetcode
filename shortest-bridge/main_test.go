package shortest_bridge

import "testing"

func Test_shortestBridge(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				grid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				grid: [][]int{
					{0, 1, 0},
					{0, 0, 0},
					{0, 0, 1},
				},
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 1},
					{1, 0, 0, 0, 1},
					{1, 0, 1, 0, 1},
					{1, 0, 0, 0, 1},
					{1, 1, 1, 1, 1},
				},
			},
			want: 1,
		},
		{
			name: "ex4",
			args: args{
				grid: [][]int{
					{0, 0, 1, 0, 1},
					{0, 1, 1, 0, 1},
					{0, 1, 0, 0, 1},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestBridge(tt.args.grid); got != tt.want {
				t.Errorf("shortestBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}
