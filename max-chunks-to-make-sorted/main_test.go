package max_chunks_to_make_sorted

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				arr: []int{4, 3, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				arr: []int{1, 0, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "ex3",
			args: args{
				arr: []int{0, 1, 3, 2, 6, 4, 7, 8, 5},
			},
			want: 4,
		},
		{
			name: "ex4",
			args: args{
				arr: []int{0, 1, 3, 2, 5, 4, 7, 8, 6},
			},
			want: 5,
		},
		{
			name: "ex5",
			args: args{
				arr: []int{1, 4, 0, 2, 3, 5},
			},
			want: 2,
		},
		{
			name: "ex6",
			args: args{
				arr: []int{1, 4, 3, 6, 0, 7, 8, 2, 5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(tt.args.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
