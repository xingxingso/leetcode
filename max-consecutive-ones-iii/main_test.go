package max_consecutive_ones_iii

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    2,
			},
			want: 6,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    3,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}

			if got := longestOnesP1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnesP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
