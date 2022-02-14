package find_duplicate_subtrees

import (
	"reflect"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3, 4, 0, 2, 4, 0, 0, 4}),
			},
			want: []*TreeNode{CreateTreeByArray([]int{4}), CreateTreeByArray([]int{2, 4})},
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{2, 1, 11, 11, 0, 1}),
			},
			want: []*TreeNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reflect.DeepEqual 这个要求数组内顺序相同 暂时不管了
			if got := findDuplicateSubtrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", got, tt.want)
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
