package binary_tree_preorder_traversal

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "equal0",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 2, 3}),
			},
			want: []int{1, 2, 3},
		},
		{
			name: "equal1",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: []int{},
		},
		{
			name: "equal2",
			args: args{
				root: CreateTreeByArray([]int{1}),
			},
			want: []int{1},
		},
		{
			name: "equal3",
			args: args{
				root: CreateTreeByArray([]int{1, 2}),
			},
			want: []int{1, 2},
		},
		{
			name: "equal4",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 2}),
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}

			if got := preorderTraversalS2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}

			if got := preorderTraversalO2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}

			if got := preorderTraversalO3(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
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
