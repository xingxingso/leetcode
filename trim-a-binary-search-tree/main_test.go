package trim_a_binary_search_tree

import (
	"reflect"
	"testing"
)

func Test_trimBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 2}),
				low:  1,
				high: 2,
			},
			want: CreateTreeByArray([]int{1, -9999, 2}),
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{3, 0, 4, -9999, 2, -9999, -9999, 1}),
				low:  1,
				high: 3,
			},
			want: CreateTreeByArray([]int{3, 2, -9999, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

// CreateTreeByArray
// 通过数组生成二叉树，nilNum 为 nil, 类似 leetcode 中数组转化为 二叉树
// [1,2,3,nilNum,4,nilNum,5,nilNum,6] 为二叉树
//                  1
//          2              3
//   nil         4      nil   5
// nil nil   nil   6
func CreateTreeByArray(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	nilNum := -9999
	queue := make([]*TreeNode, 0)
	head := &TreeNode{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums); {
		if nums[i] != nilNum {
			queue[0].Left = &TreeNode{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if i+1 < len(nums) && nums[i+1] != nilNum {
			queue[0].Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, queue[0].Right)
		}

		queue = queue[1:]
		i += 2
	}
	return head
}
