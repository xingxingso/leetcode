package closest_binary_search_tree_value

import "testing"

func Test_closestValue(t *testing.T) {
	type args struct {
		root   *TreeNode
		target float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				root:   CreateTreeByArray([]int{4, 2, 5, 1, 3}),
				target: 3.714286,
			},
			want: 4,
		},
		{
			name: "equal",
			args: args{
				root:   CreateTreeByArray([]int{5, 2, 6, 1, 4}),
				target: 3.714286,
			},
			want: 4,
		},

		{
			name: "equal",
			args: args{
				root:   CreateTreeByArray([]int{5, 2, 6, 1, 4}),
				target: 2.714286,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestValue(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("closestValue() = %v, want %v", got, tt.want)
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
