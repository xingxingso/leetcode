package longest_palindromic_substring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want2 string
	}{
		{
			name:  "equal",
			args:  args{s: "babad"},
			want:  "bab", // or "aba"
			want2: "aba",
		},
		{
			name: "equal",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "equal",
			args: args{s: "a"},
			want: "a",
		},
		{
			name:  "equal",
			args:  args{s: "ac"},
			want:  "a", // or "c"
			want2: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeO1(tt.args.s); got != tt.want && got != tt.want2 {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}

			if got := longestPalindromeO2(tt.args.s); got != tt.want && got != tt.want2 {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}

			if got := longestPalindromeO3(tt.args.s); got != tt.want && got != tt.want2 {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
