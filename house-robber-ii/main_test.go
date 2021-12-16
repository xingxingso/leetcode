package house_robber_ii

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex0",
			args: args{
				nums: []int{2, 3, 2},
			},
			want: 3,
		},
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{2, 1, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
