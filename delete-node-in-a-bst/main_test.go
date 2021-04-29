package delete_node_in_a_bst

import (
	"reflect"
	"testing"
)

type args struct {
	root *TreeNode
	key  int
}

func getTests() []struct {
	name string
	args args
	want *TreeNode
} {
	return []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{5, 3, 6, 2, 4, 0, 7}),
				key:  3,
			},
			want: CreateTreeByArray([]int{5, 4, 6, 2, 0, 0, 7}),
			// 或者 5,2,6,0,4,0,7
			// want: CreateTreeByArray([]int{5, 2, 6, 0, 4, 0, 7}),
		},
	}
}

func Test_deleteNode(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteNodeP1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNodeP1(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNodeP1() = %v, want %v", got, tt.want)
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
