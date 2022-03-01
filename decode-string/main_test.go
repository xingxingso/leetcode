package decode_string

import "testing"

func Test_decodeString(t *testing.T) {
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
				s: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "ex2",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "ex3",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
		{
			name: "ex4",
			args: args{
				s: "abc3[cd]xyz",
			},
			want: "abccdcdcdxyz",
		},
		{
			name: "ex5",
			args: args{
				s: "10[leetcode]",
			},
			want: "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := decodeStringS1(tt.args.s); got != tt.want {
			//	t.Errorf("decodeStringS1() = %v, want %v", got, tt.want)
			//}

			if got := decodeStringS2(tt.args.s); got != tt.want {
				t.Errorf("decodeStringS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
