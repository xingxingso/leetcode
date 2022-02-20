package is_graph_bipartite

import "testing"

func Test_isBipartite(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				graph: [][]int{
					{1, 2, 3},
					{0, 2},
					{0, 1, 3},
					{0, 2},
				},
			},
			want: false,
		},
		{
			name: "ex2",
			args: args{
				graph: [][]int{
					{1, 3},
					{0, 2},
					{1, 3},
					{0, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
