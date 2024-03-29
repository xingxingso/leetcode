package diameter_of_binary_tree

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
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
				root: CreateTreeByArray([]int{1, 2, 3, 4, 5}),
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3, 0, 5}),
			},
			want: 3,
		},
		{
			name: "ex3",
			args: args{
				root: CreateTreeByArray([]int{1}),
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTreeS1(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}

			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
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
