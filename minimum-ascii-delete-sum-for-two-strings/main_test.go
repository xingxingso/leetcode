package minimum_ascii_delete_sum_for_two_strings

import "testing"

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			name: "equal1",
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
