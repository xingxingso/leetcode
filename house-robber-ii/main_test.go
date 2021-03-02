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
			name: "equal0",
			args: args{
				nums: []int{2, 3, 2},
			},
			want: 3,
		},
		{
			name: "equal1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "equal2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "equal3",
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
