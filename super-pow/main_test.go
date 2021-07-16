package super_pow

import "testing"

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "ex2",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "ex3",
			args: args{
				a: 1,
				b: []int{4, 3, 3, 8, 5, 2},
			},
			want: 1,
		},
		{
			name: "ex4",
			args: args{
				a: 2147483647,
				b: []int{2, 0, 0},
			},
			want: 1198,
		},
		{
			name: "ex5",
			args: args{
				a: 1136,
				b: []int{1, 2},
			},
			want: 428,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPowP1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPowP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
