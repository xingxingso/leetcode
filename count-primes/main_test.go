package count_primes

import "testing"

func Test_countPrimes(t *testing.T) {
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
				n: 10,
			},
			want: 4,
		},
		{
			name: "ex2",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "ex3",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				n: 2,
			},
			want: 0,
		},
		{
			name: "ex5",
			args: args{
				n: 5000000,
			},
			want: 348513,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := countPrimes(tt.args.n); got != tt.want {
			//	t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			//}

			if got := countPrimesS2(tt.args.n); got != tt.want {
				t.Errorf("countPrimesS2() = %v, want %v", got, tt.want)
			}

			if got := countPrimesO1(tt.args.n); got != tt.want {
				t.Errorf("countPrimesO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
