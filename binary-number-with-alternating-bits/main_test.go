package binary_number_with_alternating_bits

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				n: 5,
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				n: 7,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				n: 11,
			},
			want: false,
		},
		{
			name: "ex4",
			args: args{
				n: 10,
			},
			want: true,
		},
		{
			name: "ex5",
			args: args{
				n: 3,
			},
			want: false,
		},
		{
			name: "ex6",
			args: args{
				n: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBits(tt.args.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
