package count_complete_tree_nodes

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3, 4, 5, 6}),
			},
			want: 6,
		},
		{
			name: "equal1",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: 0,
		},
		{
			name: "equal2",
			args: args{
				root: CreateTreeByArray([]int{1}),
			},
			want: 1,
		},
		{
			name: "equal3",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3, 4, 5}),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}

			if got := countNodesS1(tt.args.root); got != tt.want {
				t.Errorf("countNodesS1() = %v, want %v", got, tt.want)
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
