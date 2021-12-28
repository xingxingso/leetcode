package delete_nodes_and_return_forest

import (
	"reflect"
	"sort"
	"testing"
)

type args struct {
	root      *TreeNode
	to_delete []int
}

func getTests() []struct {
	name string
	args args
	want []*TreeNode
} {
	return []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root:      CreateTreeByArray([]int{1, 2, 3, 4, 5, 6, 7}),
				to_delete: []int{3, 5},
			},
			want: []*TreeNode{CreateTreeByArray([]int{1, 2, 0, 4}), CreateTreeByArray([]int{6}), CreateTreeByArray([]int{7})},
		},
		{
			name: "ex2",
			args: args{
				root:      CreateTreeByArray([]int{1, 2, 4, 0, 3}),
				to_delete: []int{3},
			},
			want: []*TreeNode{CreateTreeByArray([]int{1, 2, 4})},
		},
		{
			name: "ex3",
			args: args{
				root:      CreateTreeByArray([]int{1, 2, 3, 0, 0, 0, 4}),
				to_delete: []int{2, 1},
			},
			want: []*TreeNode{CreateTreeByArray([]int{3, 0, 4})},
		},
	}
}

func Test_delNodes(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			// 为了测试方便这里将结果排序
			got := delNodes(tt.args.root, tt.args.to_delete)
			sort.Slice(got, func(i, j int) bool {
				return got[i].Val < got[j].Val
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
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
