package reachable_nodes_in_subdivided_graph

import "testing"

func Test_reachableNodes(t *testing.T) {
	type args struct {
		edges    [][]int
		maxMoves int
		n        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				edges: [][]int{
					{0, 1, 10},
					{0, 2, 1},
					{1, 2, 2},
				},
				maxMoves: 6,
				n:        3,
			},
			want: 13,
		},
		{
			name: "ex2",
			args: args{
				edges: [][]int{
					{0, 1, 4},
					{1, 2, 6},
					{0, 2, 8},
					{1, 3, 1},
				},
				maxMoves: 10,
				n:        4,
			},
			want: 23,
		},
		{
			name: "ex3",
			args: args{
				edges: [][]int{
					{1, 2, 4},
					{1, 4, 5},
					{1, 3, 1},
					{2, 3, 4},
					{3, 4, 5},
				},
				maxMoves: 17,
				n:        5,
			},
			want: 1,
		},
		{
			name: "ex4",
			args: args{
				edges: [][]int{
					{2, 4, 2},
					{3, 4, 5},
					{2, 3, 1},
					{0, 2, 1},
					{0, 3, 5},
				},
				maxMoves: 14,
				n:        5,
			},
			want: 18,
		},
		{
			name: "ex5",
			args: args{
				edges: [][]int{
					{0, 3, 8},
					{0, 1, 4},
					{2, 4, 3},
					{1, 2, 0},
					{1, 3, 9},
					{0, 4, 7},
					{3, 4, 9},
					{1, 4, 4},
					{0, 2, 7},
					{2, 3, 1},
				},
				maxMoves: 8,
				n:        5,
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachableNodesO1(tt.args.edges, tt.args.maxMoves, tt.args.n); got != tt.want {
				t.Errorf("reachableNodesO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
