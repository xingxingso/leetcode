package longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "ex4",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}

			if got := lengthOfLongestSubstringS2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
