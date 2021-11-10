package valid_palindrome_ii

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				s: "aba",
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				s: "abac",
			},
			want: true,
		},
		{
			name: "ex3",
			args: args{
				s: "abc",
			},
			want: false,
		},
		{
			name: "ex4",
			args: args{
				s: "abcdeeca",
			},
			want: false,
		},
		{
			name: "ex5",
			args: args{
				s: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
