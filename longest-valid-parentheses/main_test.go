package longest_valid_parentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
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
				s: "(()",
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "ex3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
		{
			name: "ex5",
			args: args{
				s: ")()(",
			},
			want: 2,
		},
		{
			name: "ex6",
			args: args{
				s: "((()))())",
			},
			want: 8,
		},
		{
			name: "ex7",
			args: args{
				s: "(())())",
			},
			want: 6,
		},
		{
			name: "ex8",
			args: args{
				s: "(())()",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesO1(tt.args.s); got != tt.want {
				t.Errorf("longestValidParenthesesO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
