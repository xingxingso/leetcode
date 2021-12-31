package remove_k_digits

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
		{
			name: "ex2",
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
		{
			name: "ex3",
			args: args{
				num: "10",
				k:   2,
			},
			want: "0",
		},
		{
			name: "ex4",
			args: args{
				num: "10200",
				k:   2,
			},
			want: "0",
		},
		{
			name: "ex5",
			args: args{
				num: "34259",
				k:   3,
			},
			want: "25",
		},
		{
			name: "ex6",
			args: args{
				num: "4563",
				k:   2,
			},
			want: "43",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
