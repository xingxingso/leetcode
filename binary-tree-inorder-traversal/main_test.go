package binary_tree_inorder_traversal

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex0",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 2, 3}),
			},
			want: []int{1, 3, 2},
		},
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: []int{},
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{1}),
			},
			want: []int{1},
		},
		{
			name: "ex3",
			args: args{
				root: CreateTreeByArray([]int{1, 2}),
			},
			want: []int{2, 1},
		},
		{
			name: "ex4",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 2}),
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}

			if got := inorderTraversalS2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}

			if got := inorderTraversalO3(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
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
