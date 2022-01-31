package longest_palindromic_subsequence

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				s: "bbbab",
			},
			want: 4, // bbbb
		},
		{
			name: "ex2",
			args: args{
				s: "cbbd",
			},
			want: 2, // bb
		},
		{
			name: "ex3",
			args: args{
				s: "a",
			},
			want: 1, // a
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
