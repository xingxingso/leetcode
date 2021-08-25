package split_array_into_consecutive_subsequences

import "testing"

func Test_isPossible(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 4, 5, 5}, // 1，2，3，4，5 + 3，4，5
			},
			want: true,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{1, 2, 3, 4, 4, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.nums); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
