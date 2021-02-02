package boundary_of_binary_tree

import (
	"reflect"
	"testing"
)

func Test_boundaryOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 2, 3, 4}),
			},
			want: []int{1, 3, 4, 2},
		},
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3, 4, 5, 6, 0, 0, 0, 7, 8, 9, 10}),
			},
			want: []int{1, 2, 4, 7, 8, 9, 10, 6, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boundaryOfBinaryTreeO1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boundaryOfBinaryTree() = %v, want %v", got, tt.want)
			}

			if got := boundaryOfBinaryTreeO2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boundaryOfBinaryTree() = %v, want %v", got, tt.want)
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
