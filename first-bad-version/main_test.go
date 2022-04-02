package first_bad_version

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n   int
		bad int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				n:   5,
				bad: 4,
			},
			want: 4,
		},
		{
			name: "ex2",
			args: args{
				n:   1,
				bad: 1,
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				n:   100,
				bad: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bad = tt.args.bad
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
