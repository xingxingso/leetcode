package maximum_sum_circular_subarray

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
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
				nums: []int{1, -2, 3, -2},
			},
			want: 3, // [3]
		},
		{
			name: "ex2",
			args: args{
				nums: []int{5, -3, 5},
			},
			want: 10, // [5, 5]
		},
		{
			name: "ex3",
			args: args{
				nums: []int{3, -2, 2, -3},
			},
			want: 3, // [3] || [3, -2, 2]
		},
		{
			name: "ex4",
			args: args{
				nums: []int{-1, -2, -3, -2},
			},
			want: -1,
		},
		{
			name: "ex5",
			args: args{
				nums: []int{-2},
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircularS1(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircularS1() = %v, want %v", got, tt.want)
			}

			if got := maxSubarraySumCircularS2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircularS2() = %v, want %v", got, tt.want)
			}

			if got := maxSubarraySumCircularO1(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircularO1() = %v, want %v", got, tt.want)
			}

			if got := maxSubarraySumCircularO2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircularO2() = %v, want %v", got, tt.want)
			}

			if got := maxSubarraySumCircularO3(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircularO3() = %v, want %v", got, tt.want)
			}

			if got := maxSubarraySumCircularO4(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircularO4() = %v, want %v", got, tt.want)
			}
		})
	}
}
