package median_of_two_sorted_arrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ex1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.00000,
		},
		{
			name: "ex2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.50000,
		},
		{
			name: "ex3",
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			want: 0.00000,
		},
		{
			name: "ex4",
			args: args{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1.00000,
		},
		{
			name: "ex5",
			args: args{
				nums1: []int{2},
				nums2: []int{},
			},
			want: 2.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArraysO1(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("nums1: %v, nums2: %v, findMedianSortedArrays() = %v, want %v", tt.args.nums1, tt.args.nums2, got, tt.want)
			}
			if got := findMedianSortedArraysO2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("nums1: %v, nums2: %v, findMedianSortedArrays() = %v, want %v", tt.args.nums1, tt.args.nums2, got, tt.want)
			}
		})
	}
}
