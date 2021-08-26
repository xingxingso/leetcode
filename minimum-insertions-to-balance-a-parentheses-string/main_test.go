package minimum_insertions_to_balance_a_parentheses_string

import "testing"

func Test_minInsertions(t *testing.T) {
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
				s: "(()))",
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				s: "())",
			},
			want: 0,
		},
		{
			name: "ex3",
			args: args{
				s: "))())(",
			},
			want: 3,
		},
		{
			name: "ex4",
			args: args{
				s: "((((((",
			},
			want: 12,
		},
		{
			name: "ex5",
			args: args{
				s: ")))))))",
			},
			want: 5,
		},
		{
			name: "ex6",
			args: args{
				s: "(()))(()))()())))",
			},
			want: 4,
		},
		{
			name: "ex7",
			args: args{
				s: "(()((()((",
			},
			want: 12,
		},
		{
			name: "ex8",
			args: args{
				s: "(((()(()((())))(((()())))()())))(((()(()()((()()))",
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}

			if got := minInsertionsP1(tt.args.s); got != tt.want {
				t.Errorf("minInsertionsP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
