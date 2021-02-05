package binary_tree_vertical_order_traversal

import (
	"reflect"
	"testing"
)

func Test_verticalOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: [][]int{{9}, {3, 15}, {20}, {7}},
		},
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{3, 9, 8, 4, 6, 1, 7}),
			},
			want: [][]int{{4}, {9}, {3, 6, 1}, {8}, {7}},
		},
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{3, 9, 8, 4, 6, 1, 7, 0, 0, 0, 2, 5}),
			},
			want: [][]int{{4}, {9, 5}, {3, 6, 1}, {8, 2}, {7}},
		},
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalOrder() = %v, want %v", got, tt.want)
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
