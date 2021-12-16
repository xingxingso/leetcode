package maximum_length_of_pair_chain

import "testing"

func Test_findLongestChain(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				pairs: [][]int{{1, 2}, {2, 3}, {3, 4}},
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				pairs: [][]int{{3, 4}, {2, 4}, {1, 2}},
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				pairs: [][]int{{9, 10}, {9, 10}, {4, 5}, {-9, -3}, {-9, 1}, {0, 3}, {6, 10}, {-5, -4}, {-7, -6}},
			},
			want: 5,
		},
		{
			name: "ex4",
			args: args{
				pairs: [][]int{{7, 9}, {4, 5}, {7, 9}, {-7, -1}, {0, 10}, {3, 10}, {3, 6}, {2, 3}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestChain(tt.args.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
