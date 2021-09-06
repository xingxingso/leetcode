package assign_cookies

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				g: []int{1, 2, 3},
				s: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				g: []int{1, 2},
				s: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				g: []int{5, 3, 2, 9, 4, 6},
				s: []int{8, 5, 3, 2, 4, 5, 7, 1, 1},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
