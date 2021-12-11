package ones_and_zeroes

import "testing"

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    5,
				n:    3,
			},
			want: 4,
		},
		{
			name: "ex2",
			args: args{
				strs: []string{"10", "0", "1"},
				m:    1,
				n:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}

			if got := findMaxFormS2(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxFormS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
