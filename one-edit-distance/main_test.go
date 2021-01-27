package one_edit_distance

import "testing"

func Test_isOneEditDistance(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				s: "ab",
				t: "acb",
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				s: "cab",
				t: "ad",
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				s: "1203",
				t: "1213",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneEditDistance(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isOneEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
