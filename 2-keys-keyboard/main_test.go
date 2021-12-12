package _2_keys_keyboard

import "testing"

func Test_minSteps(t *testing.T) {
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
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				n: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStepsO1(tt.args.n); got != tt.want {
				t.Errorf("minStepsO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
