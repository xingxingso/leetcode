package minimum_moves_to_equal_array_elements_ii

import "testing"

func Test_minMoves2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{10, 3, 5},
			},
			want: 7,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{1, 5, 8, 12, 37, 23, 25, 78, 46},
			},
			want: 160,
		},
		{
			name: "ex4",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.args.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
