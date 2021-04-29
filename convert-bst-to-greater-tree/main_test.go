package convert_bst_to_greater_tree

import (
	"reflect"
	"testing"
)

type args struct {
	root *TreeNode
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
			name: "equal0",
			args: args{
				root: CreateTreeByArray([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}),
			},
			want: CreateTreeByArray([]int{30, 36, 21, 36, 35, 26, 15, -1, -1, -1, 33, -1, -1, -1, 8}),
		},
		{
			name: "equal1",
			args: args{
				root: CreateTreeByArray([]int{0, -1, 1}),
			},
			want: CreateTreeByArray([]int{1, -1, 1}),
		},
		{
			name: "equal2",
			args: args{
				root: CreateTreeByArray([]int{1, 0, 2}),
			},
			want: CreateTreeByArray([]int{3, 3, 2}),
		},
		{
			name: "equal3",
			args: args{
				root: CreateTreeByArray([]int{3, 2, 4, 1}),
			},
			want: CreateTreeByArray([]int{7, 9, 4, 10}),
		},
	}
}

func Test_convertBST(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_convertBSTO1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBSTO1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBSTO1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 这里 -1 表示 树中的空节点
func CreateTreeByArray(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	queue := make([]*TreeNode, 0)
	head := &TreeNode{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums); {
		if nums[i] != -1 {
			queue[0].Left = &TreeNode{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if i+1 < len(nums) && nums[i+1] != -1 {
			queue[0].Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, queue[0].Right)
		}

		queue = queue[1:]
		i += 2
	}
	return head
}
