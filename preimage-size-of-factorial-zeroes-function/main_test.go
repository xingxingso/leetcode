package preimage_size_of_factorial_zeroes_function

import "testing"

func Test_preimageSizeFZF(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				k: 0,
			},
			want: 5,
		},
		{
			name: "ex2",
			args: args{
				k: 5,
			},
			want: 0,
		},
		{
			name: "ex3",
			args: args{
				k: 1000,
			},
			want: 5,
		},
		{
			name: "ex4",
			args: args{
				k: 66550376,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preimageSizeFZF(tt.args.k); got != tt.want {
				t.Errorf("preimageSizeFZF() = %v, want %v", got, tt.want)
			}
		})
	}
}
