package max_area_of_island

import "testing"

type args struct {
	grid [][]int
}

func getTests() []struct {
	name string
	args args
	want int
} {
	return []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
				},
			},
			want: 6,
		},
		{
			name: "ex2",
			args: args{
				grid: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "ex3",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 1, 1},
					{0, 0, 0, 1, 1},
				},
			},
			want: 4,
		},
		{
			name: "ex4",
			args: args{
				grid: [][]int{
					{1},
				},
			},
			want: 1,
		},
	}
}

func Test_maxAreaOfIsland(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIsland(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxAreaOfIslandP1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIslandP1(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIslandP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
