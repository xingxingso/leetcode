package rearrange_string_k_distance_apart

import "testing"

/**
该测试中预期结果可能仅是正确结果的中一种
*/
func Test_rearrangeString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "equal",
			args: args{
				s: "aabbcc",
				k: 3,
			},
			want: "abcabc",
		},
		{
			name: "equal",
			args: args{
				s: "aaabc",
				k: 3,
			},
			want: "",
		},
		{
			name: "equal",
			args: args{
				s: "aaadbbcc",
				k: 2,
			},
			// want: "abacabcd",
			want: "abcadbca",
			// want: "abacacbd",
		},
		{
			name: "equal",
			args: args{
				s: "aabbccddee",
				k: 2,
			},
			want: "abcdeabcde",
			// want: "abcdecedab",
		},
		{
			name: "equal",
			args: args{
				s: "aabbccddee",
				k: 4,
			},
			// want: "abcdeabced",
			want: "abcdeabcde",
		},
		{
			name: "equal",
			args: args{
				s: "aabbccddee",
				k: 5,
			},
			want: "abcdeabcde",
		},
		{
			name: "equal",
			args: args{
				s: "a",
				k: 0,
			},
			want: "a",
		},
		{
			name: "equal",
			args: args{
				s: "a",
				k: 2,
			},
			want: "a",
		},
		{
			name: "equal",
			args: args{
				s: "aa",
				k: 2,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("rearrangeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
