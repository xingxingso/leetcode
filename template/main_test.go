package template

import (
	"reflect"
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
			want: 2.00000,
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

func stringSliceSliceEqual(v1, v2 [][]string) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if reflect.DeepEqual(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	return true
}

func intSliceSliceEqual(v1, v2 [][]int) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if len(vv1) != len(vv2) {
				continue
			}
			if sliceSameValue(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	return true
}

func sliceSameValue(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
loop:
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				continue loop
			}
		}
		return false
	}

	return true
}

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func root1() *TreeNode {
	node5 := &TreeNode{}
	node5.Val = 5

	node4 := &TreeNode{}
	node4.Val = 4

	node3 := &TreeNode{}
	node3.Val = 3

	node2 := &TreeNode{}
	node2.Val = 2
	node2.Left = node4
	node2.Right = node5

	node1 := &TreeNode{}
	node1.Val = 1
	node1.Left = node2
	node1.Right = node3

	return node1
}

func want1() *TreeNode {
	node3 := &TreeNode{}
	node3.Val = 3

	node1 := &TreeNode{}
	node1.Val = 1

	node5 := &TreeNode{}
	node5.Val = 5

	node2 := &TreeNode{}
	node2.Val = 2
	node2.Left = node3
	node2.Right = node1

	node4 := &TreeNode{}
	node4.Val = 4
	node4.Left = node5
	node4.Right = node2

	return node4
}
