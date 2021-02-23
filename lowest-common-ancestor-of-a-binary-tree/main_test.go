package lowest_common_ancestor_of_a_binary_tree

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	root1 := CreateTreeByArray([]int{3, 5, 1, 6, 2, 9, 8, 0, 0, 7, 4})
	root2 := CreateTreeByArray([]int{3, 5, 1, 6, 2, 9, 8, 0, 0, 7, 4})
	root3 := CreateTreeByArray([]int{1, 2})

	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "equal0",
			args: args{
				root: root1,
				p:    root1.Left,  // 5
				q:    root1.Right, // 1
			},
			want: root1, // 3
		},
		{
			name: "equal1",
			args: args{
				root: root2,
				p:    root2.Left,             // 5
				q:    root2.Left.Right.Right, // 4
			},
			want: root2.Left, // 5
		},
		{
			name: "equal2",
			args: args{
				root: root3,
				p:    root3,      // 1
				q:    root3.Left, // 2
			},
			want: root3, // 1
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}

			if got := lowestCommonAncestorO1(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
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
