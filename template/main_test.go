package template

import (
	"testing"
)

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
			name: "equal",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 0.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("nums1: %v, nums2: %v, findMedianSortedArrays() = %v, want %v", tt.args.nums1, tt.args.nums2, got, tt.want)
			}
		})
	}
}
