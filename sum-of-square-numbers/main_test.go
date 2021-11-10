package sum_of_square_numbers

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				c: 5,
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				c: 3,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				c: 4,
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				c: 2,
			},
			want: true,
		},
		{
			name: "ex5",
			args: args{
				c: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}

			if got := judgeSquareSumO1(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSumO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
