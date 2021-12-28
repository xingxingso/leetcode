package longest_duplicate_substring

import "testing"

func Test_longestDupSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				s: "banana",
			},
			want: "ana",
		},
		{
			name: "ex2",
			args: args{
				s: "abcd",
			},
			want: "",
		},
		{
			name: "ex3",
			args: args{
				s: "aa",
			},
			want: "a",
		},
		{
			name: "ex4",
			args: args{
				s: "aba",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDupSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
