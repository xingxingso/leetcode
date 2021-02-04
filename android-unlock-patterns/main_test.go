package android_unlock_patterns

import "testing"

func Test_numberOfPatterns(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				m: 1,
				n: 1,
			},
			want: 9,
		},
		{
			name: "equal",
			args: args{
				m: 1,
				n: 2,
			},
			want: 65,
		},
		{
			name: "equal",
			args: args{
				m: 4,
				n: 7,
			},
			want: 107704,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPatterns(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("numberOfPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}
