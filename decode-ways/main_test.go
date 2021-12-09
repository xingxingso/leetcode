package decode_ways

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				s: "12",
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				s: "226",
			},
			want: 3,
		},
		{
			name: "ex3",
			args: args{
				s: "0",
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				s: "06",
			},
			want: 0,
		},
		{
			name: "ex5",
			args: args{
				s: "2101",
			},
			want: 1,
		},
		{
			name: "ex6",
			args: args{
				s: "10",
			},
			want: 1,
		},
		{
			name: "ex7",
			args: args{
				s: "1123", // 1 1 2 3|11 2 3| 1 12 3|1 1 23|11 23|
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
