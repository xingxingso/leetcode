package factorial_trailing_zeroes

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				n: 3,
			},
			want: 0,
		},
		{
			name: "ex2",
			args: args{
				n: 5,
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				n: 100,
			},
			want: 24,
		},
		{
			name: "ex4",
			args: args{
				n: 7, // 7! = 5040
			},
			want: 1, // 求尾数中的 0
		},
		{
			name: "ex5",
			args: args{
				n: 9,
			},
			want: 1,
		},
		{
			name: "ex6",
			args: args{
				n: 24,
			},
			want: 4,
		},
		{
			name: "ex7",
			args: args{
				n: 25,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
