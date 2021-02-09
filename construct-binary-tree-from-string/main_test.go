package construct_binary_tree_from_string

import (
	"testing"
)

func Test_str2tree(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "equal0",
			args: args{
				s: "4(2(3)(1))(6(5))",
			},
			want: CreateTreeByArray([]int{4, 2, 6, 3, 1, 5}),
		},
		{
			name: "equal1",
			args: args{
				s: "8",
			},
			want: CreateTreeByArray([]int{8}),
		},
		{
			name: "equal2",
			args: args{
				s: "-8",
			},
			want: CreateTreeByArray([]int{-8}),
		},
		{
			name: "equal3",
			args: args{
				s: "",
			},
			want: CreateTreeByArray([]int{}),
		},
		{
			name: "equal0",
			args: args{
				s: "-4(25(31)(-128))(6(-5))",
			},
			want: CreateTreeByArray([]int{-4, 25, 6, 31, -128, -5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := str2tree(tt.args.s); !isSameTree(got, tt.want) {
				t.Errorf("str2tree() = %v, want %v", got, tt.want)
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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
