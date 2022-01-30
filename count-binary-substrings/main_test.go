package count_binary_substrings

import "testing"

func Test_countBinarySubstrings(t *testing.T) {
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
				s: "00110011",
			},
			want: 6, // "0011"、"01"、"1100"、"10"、"0011" 和 "01"
		},
		{
			name: "ex2",
			args: args{
				s: "10101",
			},
			want: 4, // "10"、"01"、"10"、"01"
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings(tt.args.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
