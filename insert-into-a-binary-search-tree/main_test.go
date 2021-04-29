package insert_into_a_binary_search_tree

import (
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{4, 2, 7, 1, 3}),
				val:  5,
			},
			want: CreateTreeByArray([]int{4, 2, 7, 1, 3, 5}), // 或者 5 2 7 1 3 0 0 0 0 0 4
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{40, 20, 60, 10, 30, 50, 70}),
				val:  25,
			},
			want: CreateTreeByArray([]int{40, 20, 60, 10, 30, 50, 70, 0, 0, 25}),
		},
		{
			name: "ex3",
			args: args{
				root: CreateTreeByArray([]int{4, 2, 7, 1, 3, 0, 0, 0, 0, 0, 0}),
				val:  5,
			},
			want: CreateTreeByArray([]int{4, 2, 7, 1, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
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
