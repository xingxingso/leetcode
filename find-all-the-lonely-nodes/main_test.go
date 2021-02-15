package find_all_the_lonely_nodes

import (
	"reflect"
	"testing"
)

func Test_getLonelyNodes(t *testing.T) {
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
				root: CreateTreeByArray([]int{1, 2, 3, 0, 4}),
			},
			want: []int{4},
		},
		{
			name: "equal1",
			args: args{
				root: CreateTreeByArray([]int{7, 1, 4, 6, 0, 5, 3, 0, 0, 0, 0, 0, 2}),
			},
			want: []int{6, 2},
		},
		{
			name: "equal2",
			args: args{
				root: CreateTreeByArray([]int{11, 99, 88, 77, 0, 0, 66, 55, 0, 0, 44, 33, 0, 0, 22}),
			},
			want: []int{77, 55, 33, 66, 44, 22},
		},
		{
			name: "equal3",
			args: args{
				root: CreateTreeByArray([]int{197}),
			},
			want: []int{},
		},
		{
			name: "equal4",
			args: args{
				root: CreateTreeByArray([]int{31, 0, 78, 0, 28}),
			},
			want: []int{78, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLonelyNodes(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLonelyNodes() = %v, want %v", got, tt.want)
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
