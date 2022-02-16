package minimum_absolute_difference_in_bst

import "testing"

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{4, 2, 6, 1, 3}),
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 48, -9999, -9999, 12, 49}),
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				root: CreateTreeByArray([]int{236, 104, 701, -9999, 227, -9999, 911}),
			},
			want: 9,
		},
		{
			name: "ex4",
			args: args{
				root: CreateTreeByArray([]int{543, 384, 652, -9999, 445, -9999, 699}),
			},
			want: 47,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
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
