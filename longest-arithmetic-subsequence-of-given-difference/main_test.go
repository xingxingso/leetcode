package longest_arithmetic_subsequence_of_given_difference

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		arr        []int
		difference int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				arr:        []int{1, 2, 3, 4},
				difference: 1,
			},
			want: 4,
		},
		{
			name: "ex2",
			args: args{
				arr:        []int{1, 3, 5, 7},
				difference: 1,
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				arr:        []int{1, 5, 7, 8, 5, 3, 4, 2, 1},
				difference: -2,
			},
			want: 4,
		},
		{
			name: "ex4",
			args: args{
				arr:        []int{3, 4, -3, -2, -4},
				difference: -5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequenceS1(tt.args.arr, tt.args.difference); got != tt.want {
				t.Errorf("longestSubsequenceS1() = %v, want %v", got, tt.want)
			}

			if got := longestSubsequenceS2(tt.args.arr, tt.args.difference); got != tt.want {
				t.Errorf("longestSubsequenceS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
