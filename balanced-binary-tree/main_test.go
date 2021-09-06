package balanced_binary_tree

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 2, 3, 3, 0, 0, 4, 4}),
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 2, 3, 0, 0, 3, 4, 0, 0, 4}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalancedO1(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}

			if got := isBalancedO2(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func CreateTreeByArray(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	queue := make([]*TreeNode, 0)
	head := &TreeNode{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums)-1; {
		if nums[i] != 0 {
			queue[0].Left = &TreeNode{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if nums[i+1] != 0 {
			queue[0].Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, queue[0].Right)
		}

		queue = queue[1:]
		i += 2
	}
	return head
}
