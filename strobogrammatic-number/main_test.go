package strobogrammatic_number

import "testing"

func Test_isStrobogrammatic(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				num: "69",
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				num: "88",
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				num: "962",
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				num: "1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStrobogrammatic(tt.args.num); got != tt.want {
				t.Errorf("isStrobogrammatic() = %v, want %v", got, tt.want)
			}
		})
	}
}
