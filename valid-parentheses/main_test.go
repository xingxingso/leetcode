package valid_parentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal0",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "equal1",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "equal3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "equal4",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "equal5",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
		{
			name: "equal6",
			args: args{
				s: "(){}}{",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
