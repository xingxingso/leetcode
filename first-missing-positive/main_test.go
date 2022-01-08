package first_missing_positive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		},
		{
			name: "ex4",
			args: args{
				nums: []int{3, 2, 5, 4, 1},
			},
			want: 6,
		},
		{
			name: "ex5",
			args: args{
				nums: []int{3, 2, 5, 4, 1},
			},
			want: 6,
		},
		{
			name: "ex6",
			args: args{
				nums: []int{0},
			},
			want: 1,
		},
		{
			name: "ex7",
			args: args{
				nums: []int{1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
