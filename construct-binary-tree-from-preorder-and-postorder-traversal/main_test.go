package construct_binary_tree_from_preorder_and_postorder_traversal

import (
	"reflect"
	"testing"
)

func Test_constructFromPrePost(t *testing.T) {
	type args struct {
		preorder  []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex1",
			args: args{
				preorder:  []int{1, 2, 4, 5, 3, 6, 7},
				postorder: []int{4, 5, 2, 6, 7, 3, 1},
			},
			want: CreateTreeByArray([]int{1, 2, 3, 4, 5, 6, 7}),
		},
		{
			name: "ex2",
			args: args{
				preorder:  []int{1},
				postorder: []int{1},
			},
			want: CreateTreeByArray([]int{1}),
		},
		{
			name: "ex3",
			args: args{
				preorder:  []int{1, 2, 3},
				postorder: []int{2, 3, 1},
			},
			want: CreateTreeByArray([]int{1, 2, 3}),
		},
		{
			name: "ex4",
			args: args{
				preorder:  []int{1, 2, 3},
				postorder: []int{3, 2, 1},
			},
			//want: CreateTreeByArray([]int{1, -9999, 2, -9999, 3}), //或者 1, -9999, 2, 3
			//want: CreateTreeByArray([]int{1, -9999, 2, 3}),
			want: CreateTreeByArray([]int{1, 2, -9999, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructFromPrePost(tt.args.preorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePost() = %v, want %v", got, tt.want)
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
