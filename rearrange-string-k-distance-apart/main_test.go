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
			name: "ex1",
			args: args{
				s: "aabbcc",
				k: 3,
			},
			want: "abcabc",
		},
		{
			name: "ex2",
			args: args{
				s: "aaabc",
				k: 3,
			},
			want: "",
		},
		{
			name: "ex3",
			args: args{
				s: "aaadbbcc",
				k: 2,
			},
			// want: "abacabcd",
			want: "abcadbca",
			// want: "abacacbd",
		},
		{
			name: "ex4",
			args: args{
				s: "aabbccddee",
				k: 2,
			},
			want: "abcdeabcde",
			// want: "abcdecedab",
		},
		{
			name: "ex5",
			args: args{
				s: "aabbccddee",
				k: 4,
			},
			// want: "abcdeabced",
			want: "abcdeabcde",
		},
		{
			name: "ex6",
			args: args{
				s: "aabbccddee",
				k: 5,
			},
			want: "abcdeabcde",
		},
		{
			name: "ex7",
			args: args{
				s: "a",
				k: 0,
			},
			want: "a",
		},
		{
			name: "ex8",
			args: args{
				s: "a",
				k: 2,
			},
			want: "a",
		},
		{
			name: "ex9",
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
