package maximum_length_of_subarray_with_positive_product

import "testing"

func Test_getMaxLen(t *testing.T) {
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
				nums: []int{1, -2, -3, 4},
			},
			want: 4,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{0, 1, -2, -3, -4},
			},
			want: 3,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{-1, -2, -3, 0, 1},
			},
			want: 2,
		},
		{
			name: "ex4",
			args: args{
				nums: []int{1, 2, 3, 5, -6, 4, 0, 10},
			},
			want: 4,
		},
		{
			name: "ex5",
			args: args{
				nums: []int{5, -20, -20, -39, -5, 0, 0, 0, 36, -32, 0, -7, -10, -7, 21, 20, -12, -34, 26, 2},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxLen(tt.args.nums); got != tt.want {
				t.Errorf("getMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
