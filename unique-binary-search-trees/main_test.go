package unique_binary_search_trees

import "testing"

func Test_numTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				n: 3,
			},
			want: 5,
		},
		{
			name: "ex2",
			args: args{
				n: 5,
			},
			want: 42,
		},
		{
			name: "ex3",
			args: args{
				n: 6,
			},
			want: 132,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
