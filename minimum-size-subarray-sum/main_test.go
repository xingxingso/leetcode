package minimum_size_subarray_sum

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
