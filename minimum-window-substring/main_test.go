package minimum_window_substring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "equal0",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "equal1",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "equal3",
			args: args{
				s: "ADOBAECODEBANC",
				t: "AA",
			},
			want: "ADOBA",
		},
		{
			name: "equal4",
			args: args{
				s: "ADBOBACODEBANC",
				t: "ABC",
			},
			want: "BAC",
		},
		{
			name: "equal5",
			args: args{
				s: "A",
				t: "ABC",
			},
			want: "",
		},
		{
			name: "equal6",
			args: args{
				s: "ABAACBAB",
				t: "ABC",
			},
			want: "ACB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}

			if got := minWindowO1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
