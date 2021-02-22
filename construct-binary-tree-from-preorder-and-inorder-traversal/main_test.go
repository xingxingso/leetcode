package construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "equal0",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: CreateTreeByArray([]int{3, 9, 20, 0, 0, 15, 7}),
		},
		{
			name: "equal1",
			args: args{
				preorder: []int{},
				inorder:  []int{},
			},
			want: CreateTreeByArray([]int{}),
		},
		{
			name: "equal2",
			args: args{
				preorder: []int{1},
				inorder:  []int{1},
			},
			want: CreateTreeByArray([]int{1}),
		},
		{
			name: "equal3",
			args: args{
				preorder: []int{3, 20, 15},
				inorder:  []int{3, 15, 20},
			},
			want: CreateTreeByArray([]int{3, 0, 20, 15}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
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
