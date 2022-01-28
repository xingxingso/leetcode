package all_paths_from_source_lead_to_destination

import "testing"

func Test_leadsToDestination(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				n:           3,
				edges:       [][]int{{0, 1}, {0, 2}},
				source:      0,
				destination: 2,
			},
			want: false,
		},
		{
			name: "ex2",
			args: args{
				n:           4,
				edges:       [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 1}},
				source:      0,
				destination: 3,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				n:           4,
				edges:       [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
				source:      0,
				destination: 3,
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				n:           3,
				edges:       [][]int{{0, 1}, {1, 1}, {1, 2}},
				source:      0,
				destination: 2,
			},
			want: false,
		},
		{
			name: "ex5",
			args: args{
				n:           2,
				edges:       [][]int{{0, 1}, {1, 1}},
				source:      0,
				destination: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leadsToDestination(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("leadsToDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
