package generate_parentheses

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ex1",
			args: args{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "ex2",
			args: args{
				n: 1,
			},
			want: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}

			if got := generateParenthesisP1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesisP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
