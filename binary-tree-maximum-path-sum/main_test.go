package binary_tree_maximum_path_sum

import "testing"

func Test_maxPathSum(t *testing.T) {
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
				root: CreateTreeByArray([]int{1, 2, 3}),
			},
			want: 6,
		},
		{
			name: "equal1",
			args: args{
				root: CreateTreeByArray([]int{-10, 9, 20, 0, 0, 15, 7}),
			},
			want: 42,
		},
		{
			name: "equal2",
			args: args{
				root: CreateTreeByArray([]int{-1, 3, 2, 4, 0, 5, 6}),
			},
			want: 14,
		},
		{
			name: "equal3",
			args: args{
				root: CreateTreeByArray([]int{-1, -3, -20, 4, 0, 5, 6}),
			},
			want: 6,
		},
		{
			name: "equal4",
			args: args{
				root: CreateTreeByArray([]int{-3}),
			},
			want: -3,
		},
		{
			name: "equal5",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
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
