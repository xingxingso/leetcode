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
			name:  "ex1",
			args:  args{s: "babad"},
			want:  "bab", // or "aba"
			want2: "aba",
		},
		{
			name: "ex2",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "ex3",
			args: args{s: "a"},
			want: "a",
		},
		{
			name:  "ex4",
			args:  args{s: "ac"},
			want:  "a", // or "c"
			want2: "c",
		},
		{
			name: "ex5",
			args: args{s: "aacabdkacaa"},
			want: "aca",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeS1(tt.args.s); got != tt.want && got != tt.want2 {
				t.Errorf("longestPalindromeS1() = %v, want %v", got, tt.want)
			}

			if got := longestPalindromeO1(tt.args.s); got != tt.want && got != tt.want2 {
				t.Errorf("longestPalindromeO1() = %v, want %v", got, tt.want)
			}

			if got := longestPalindromeO2(tt.args.s); got != tt.want && got != tt.want2 {
				t.Errorf("longestPalindromeO2() = %v, want %v", got, tt.want)
			}

			if got := longestPalindromeO3(tt.args.s); got != tt.want && got != tt.want2 {
				t.Errorf("longestPalindromeO3() = %v, want %v", got, tt.want)
			}
		})
	}
}
