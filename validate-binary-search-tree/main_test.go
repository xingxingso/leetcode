package validate_binary_search_tree

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal0",
			args: args{
				root: CreateTreeByArray([]int{2, 1, 3}),
			},
			want: true,
		},
		{
			name: "equal1",
			args: args{
				root: CreateTreeByArray([]int{5, 1, 4, 0, 0, 3, 6}),
			},
			want: false,
		},
		{
			name: "equal2",
			args: args{
				root: CreateTreeByArray([]int{4, 1, 5, 0, 2, 3, 6}),
			},
			want: false,
		},
		{
			name: "equal3",
			args: args{
				root: CreateTreeByArray([]int{4, 1, 5, 0, 2, 0, 6}),
			},
			want: true,
		},
		{
			name: "equal4",
			args: args{
				root: CreateTreeByArray([]int{1, 1}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}

			if got := isValidBSTO1(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
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
