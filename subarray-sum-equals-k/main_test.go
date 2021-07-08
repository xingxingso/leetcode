package subarray_sum_equals_k

import "testing"

func Test_subarraySum(t *testing.T) {
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1, 2, 4, -3, 2, 1},
				k:    3,
			},
			want: 4,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{1, -1, 1, 2, -3, -2, 2},
				k:    0,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySumP1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySumP1() = %v, want %v", got, tt.want)
			}

			if got := subarraySumP2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySumP2() = %v, want %v", got, tt.want)
			}
		})
	}
}
