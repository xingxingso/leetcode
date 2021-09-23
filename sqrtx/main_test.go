package sqrtx

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				x: 2000000000,
			},
			want: 44721,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}

			if got := mySqrtO1(tt.args.x); got != tt.want {
				t.Errorf("mySqrtO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
