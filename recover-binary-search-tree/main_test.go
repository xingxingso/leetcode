package recover_binary_search_tree

import (
	"reflect"
	"testing"
)

type args struct {
	root *TreeNode
}

func getTests() []struct {
	name string
	args args
	want *TreeNode
} {
	return []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{1, 3, 0, 0, 2}),
			},
			want: CreateTreeByArray([]int{3, 1, 0, 0, 2}),
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{3, 1, 4, 0, 0, 2}),
			},
			want: CreateTreeByArray([]int{2, 1, 4, 0, 0, 3}),
		},
	}
}

func Test_recoverTree(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if recoverTree(tt.args.root); !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}

// CreateTreeByArray
// 通过数组生成二叉树，0 为 nil, 类似 leetcode 中数组转化为 二叉树
// [1,2,3,0,4,0,5,0,6] 为二叉树
//                  1
//          2              3
//   nil         4      nil   5
// nil nil   nil   6
func CreateTreeByArray(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	queue := make([]*TreeNode, 0)
	head := &TreeNode{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums); {
		if nums[i] != 0 {
			queue[0].Left = &TreeNode{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if i+1 < len(nums) && nums[i+1] != 0 {
			queue[0].Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, queue[0].Right)
		}

		queue = queue[1:]
		i += 2
	}
	return head
}
