package power_of_three

import "testing"

func Test_isPowerOfThree(t *testing.T) {
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
				n: 27,
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				n: 9,
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				n: 45,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}

			if got := isPowerOfThreeS2(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThreeS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
