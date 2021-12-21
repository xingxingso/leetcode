package base_7

import "testing"

func Test_convertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				num: 100,
			},
			want: "202",
		},
		{
			name: "ex2",
			args: args{
				num: -7,
			},
			want: "-10",
		},
		{
			name: "ex3",
			args: args{
				num: 0,
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
