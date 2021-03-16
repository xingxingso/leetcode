package target_sum

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				S:    3,
			},
			want: 5,
		},
		{
			name: "equal1",
			args: args{
				nums: []int{1000},
				S:    -1000,
			},
			want: 1,
		},
		{
			name: "equal2",
			args: args{
				nums: []int{1, 2, 7, 9, 981},
				S:    1000000000,
			},
			want: 0,
		},
		{
			name: "equal3",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				S:    -10000000,
			},
			want: 0,
		},
		{
			name: "equal4",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
				S:    1,
			},
			want: 256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}

			if got := findTargetSumWaysS2(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}

			if got := findTargetSumWaysS3(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
