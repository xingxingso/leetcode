package minimum_insertion_steps_to_make_a_string_palindrome

import "testing"

func Test_minInsertions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				s: "zzazz",
			},
			want: 0,
		},
		{
			name: "equal1",
			args: args{
				s: "mbadm",
			},
			want: 2,
		},
		{
			name: "equal2",
			args: args{
				s: "leetcode",
			},
			want: 5,
		},
		{
			name: "equal3",
			args: args{
				s: "g",
			},
			want: 0,
		},
		{
			name: "equal4",
			args: args{
				s: "no",
			},
			want: 1,
		},
		{
			name: "equal5",
			args: args{
				s: "abcab", // abacaba 或 abcacba
			},
			want: 2,
		},
		{
			name: "equal6",
			args: args{
				s: "status",
			},
			want: 1,
		},
		{
			name: "equal7",
			args: args{
				s: "zjveiiwvc",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}

			if got := minInsertionsS2(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
