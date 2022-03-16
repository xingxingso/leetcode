package sum_of_two_integers

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
		{
			name: "ex3",
			args: args{
				a: -1,
				b: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}

			if got := getSumO1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSumO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
