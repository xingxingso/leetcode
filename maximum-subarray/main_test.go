package maximum_subarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex0",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "ex1",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{-100000},
			},
			want: -100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}

			if got := maxSubArrayO2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
