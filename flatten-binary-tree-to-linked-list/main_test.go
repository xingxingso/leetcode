package flatten_binary_tree_to_linked_list

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "equal0",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 5, 3, 4, 0, 6}),
			},
			want: CreateTreeByArray([]int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}),
		},
		{
			name: "equal1",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: CreateTreeByArray([]int{}),
		},
		{
			name: "equal2",
			args: args{
				root: CreateTreeByArray([]int{1}),
			},
			want: CreateTreeByArray([]int{1}),
		},
		{
			name: "equal3",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3}),
			},
			want: CreateTreeByArray([]int{1, 0, 2, 0, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if flatten(tt.args.root); !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.args.root, tt.want)
			}

			if flattenO1(tt.args.root); !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.args.root, tt.want)
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
