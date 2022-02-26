package longest_common_subsequence

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "ex3",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				text1: "ezupkr",
				text2: "ubmrapg",
			},
			want: 2,
		},
		{
			name: "ex5",
			args: args{
				text1: "bsbininm",
				text2: "jmjkbkjkv",
			},
			want: 1,
		},
		{
			name: "ex6",
			args: args{
				text1: "mhunuzqrkzsnidwbun",
				text2: "szulspmhwpazoxijwbq",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequenceS1(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceS1() = %v, want %v", got, tt.want)
			}

			if got := longestCommonSubsequenceP1(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceP1() = %v, want %v", got, tt.want)
			}

			if got := longestCommonSubsequenceP2(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceP2() = %v, want %v", got, tt.want)
			}

			if got := longestCommonSubsequenceP3(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceP3() = %v, want %v", got, tt.want)
			}
		})
	}
}
