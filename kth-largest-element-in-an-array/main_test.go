package kth_largest_element_in_an_array

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{1, 3, 5, 2, 2},
				k:    3,
			},
			want: 2,
		},
		{
			name: "ex4",
			args: args{
				nums: []int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2},
				k:    3,
			},
			want: 9,
		},
		{
			name: "ex5",
			args: args{
				nums: []int{2, 1},
				k:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestS1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestS1() = %v, want %v", got, tt.want)
			}

			//if got := findKthLargestP1(tt.args.nums, tt.args.k); got != tt.want {
			//	t.Errorf("findKthLargestP1() = %v, want %v", got, tt.want)
			//}
			//
			//if got := findKthLargestP2(tt.args.nums, tt.args.k); got != tt.want {
			//	t.Errorf("findKthLargestP2() = %v, want %v", got, tt.want)
			//}
		})
	}
}
