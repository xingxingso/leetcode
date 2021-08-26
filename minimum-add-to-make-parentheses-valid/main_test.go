package minimum_add_to_make_parentheses_valid

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
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
				s: "())",
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				s: "(((",
			},
			want: 3,
		},
		{
			name: "ex3",
			args: args{
				s: "()",
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				s: "()))((",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAddToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
