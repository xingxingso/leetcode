package maximum_length_of_repeated_subarray

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{3, 2, 1, 4, 7},
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				nums1: []int{0, 0, 0, 0, 0},
				nums2: []int{0, 0, 0, 0, 0},
			},
			want: 5,
		},
		{
			name: "ex3",
			args: args{
				nums1: []int{0, 0, 0, 0, 1},
				nums2: []int{1, 0, 0, 0, 0},
			},
			want: 4,
		},
		{
			name: "ex4",
			args: args{
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{3, 2, 1, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := findLength(tt.args.nums1, tt.args.nums2); got != tt.want {
			//	t.Errorf("findLength() = %v, want %v", got, tt.want)
			//}

			if got := findLengthO1(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findLengthO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
