package subtree_of_another_tree

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				root:    CreateTreeByArray([]int{3, 4, 5, 1, 2}),
				subRoot: CreateTreeByArray([]int{4, 1, 2}),
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				root:    CreateTreeByArray([]int{3, 4, 5, 1, 2, -9999, -9999, -9999, -9999, 0}),
				subRoot: CreateTreeByArray([]int{4, 1, 2}),
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				root:    CreateTreeByArray([]int{1, 1}),
				subRoot: CreateTreeByArray([]int{1}),
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				root:    CreateTreeByArray([]int{3, 4, 5, 1, -9999, 2}),
				subRoot: CreateTreeByArray([]int{3, 1, 2}),
			},
			want: false,
		},
		{
			name: "ex5",
			args: args{
				root:    CreateTreeByArray([]int{4, 1, -9999, 1, -9999, 6, 7}),
				subRoot: CreateTreeByArray([]int{4, 1, -9999, 6, 7}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
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
