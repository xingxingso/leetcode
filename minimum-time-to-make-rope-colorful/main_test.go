package minimum_time_to_make_rope_colorful

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		colors     string
		neededTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				colors:     "abaac",
				neededTime: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				colors:     "abc",
				neededTime: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "ex3",
			args: args{
				colors:     "aabaa",
				neededTime: []int{1, 2, 3, 4, 1},
			},
			want: 2,
		},
		{
			name: "ex4",
			args: args{
				colors:     "cddcdcae",
				neededTime: []int{4, 8, 8, 4, 4, 5, 4, 2},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.colors, tt.args.neededTime); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
