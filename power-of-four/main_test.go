package power_of_four

import "testing"

func Test_isPowerOfFour(t *testing.T) {
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
				n: 16,
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				n: 5,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				n: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
