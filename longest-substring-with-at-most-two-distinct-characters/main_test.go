package longest_substring_with_at_most_two_distinct_characters

import "testing"

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				s: "eceba",
			},
			want: 3,
		},
		{
			name: "equal",
			args: args{
				s: "ccaabbb",
			},
			want: 5,
		},
		{
			name: "equal",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "equal",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "equal",
			args: args{
				s: "aa",
			},
			want: 2,
		},
		{
			name: "equal",
			args: args{
				s: "aab",
			},
			want: 3,
		},
		{
			name: "equal",
			args: args{
				s: "abaccc",
			},
			want: 4,
		},
		{
			name: "equal",
			args: args{
				s: "abcabcabc",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringTwoDistinct(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinct() = %v, want %v", got, tt.want)
			}

			if got := lengthOfLongestSubstringTwoDistinctO1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinctO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
