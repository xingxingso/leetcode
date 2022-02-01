package basic_calculator_ii

import "testing"

func Test_calculate(t *testing.T) {
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
				s: "3+2*2",
			},
			want: 7,
		},
		{
			name: "ex2",
			args: args{
				s: " 3/2 ",
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				s: " 3+5 / 2 ",
			},
			want: 5,
		},
		{
			name: "ex4",
			args: args{
				s: "1-1+1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
