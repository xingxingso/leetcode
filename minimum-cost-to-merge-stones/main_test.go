package minimum_cost_to_merge_stones

import "testing"

func Test_mergeStones(t *testing.T) {
	type args struct {
		stones []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				stones: []int{3, 2, 4, 1},
				k:      2,
			},
			want: 20,
		},
		{
			name: "ex2",
			args: args{
				stones: []int{3, 2, 4, 1},
				k:      3,
			},
			want: -1,
		},
		{
			name: "ex3",
			args: args{
				stones: []int{3, 5, 1, 2, 6},
				k:      3,
			},
			want: 25,
		},
		{
			name: "ex4",
			args: args{
				stones: []int{6, 4, 4, 6}, // 这个证明了 简单贪心的问题， 6,8,6 -> 14,6 -> 20 = 42 > 10,10->20
				k:      2,
			},
			want: 40,
		},
		{
			name: "ex5",
			args: args{
				stones: []int{69, 39, 79, 78, 16, 6, 36, 97, 79, 27, 14, 31, 4}, // 超出时限
				k:      2,
			},
			want: 1957,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := mergeStonesS2(tt.args.stones, tt.args.k); got != tt.want {
			//	t.Errorf("mergeStonesS2() = %v, want %v", got, tt.want)
			//}

			if got := mergeStonesP1(tt.args.stones, tt.args.k); got != tt.want {
				t.Errorf("mergeStonesP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
