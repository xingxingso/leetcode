package super_ugly_number

import "testing"

func Test_nthSuperUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				n:      12,
				primes: []int{2, 7, 13, 19},
			},
			want: 32, // [1,2,4,7,8,13,14,16,19,26,28,32]
		},
		{
			name: "ex2",
			args: args{
				n:      1,
				primes: []int{2, 3, 5}, // [1,2,3,4,5]
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				n:      5,
				primes: []int{2, 3, 5}, // [1,2,3,4,5]
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthSuperUglyNumber(tt.args.n, tt.args.primes); got != tt.want {
				t.Errorf("nthSuperUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
