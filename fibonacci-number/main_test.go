package fibonacci_number

import "testing"

func Test_fib(t *testing.T) {
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
				n: 2,
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "ex4",
			args: args{
				n: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
