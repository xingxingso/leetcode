package path_sum_iii

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				root:      CreateTreeByArray([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}),
				targetSum: 8,
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				root:      CreateTreeByArray([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}),
				targetSum: 22,
			},
			want: 3,
		},
		{
			name: "ex3",
			args: args{
				root:      CreateTreeByArray([]int{}),
				targetSum: 1,
			},
			want: 0,
		},
		{
			name: "ex4",
			args: args{
				root:      CreateTreeByArray([]int{1, 2}),
				targetSum: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}

			if got := pathSumO1(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSumO1() = %v, want %v", got, tt.want)
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
