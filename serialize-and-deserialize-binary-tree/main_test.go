package serialize_and_deserialize_binary_tree

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
			name: "ex1",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3, 0, 4}),
			},
			want: CreateTreeByArray([]int{1, 2, 3, 0, 4}),
		},
		{
			name: "ex2",
			args: args{
				root: CreateTreeByArray([]int{1, 2, 3, 0, 0, 4, 5}),
			},
			want: CreateTreeByArray([]int{1, 2, 3, 0, 0, 4, 5}),
		},
		{
			name: "ex3",
			args: args{
				root: CreateTreeByArray([]int{}),
			},
			want: CreateTreeByArray([]int{}),
		},
		{
			name: "ex4",
			args: args{
				root: CreateTreeByArray([]int{1}),
			},
			want: CreateTreeByArray([]int{1}),
		},
		{
			name: "ex5",
			args: args{
				root: CreateTreeByArray([]int{1, 2}),
			},
			want: CreateTreeByArray([]int{1, 2}),
		},
	}
}

func TestCodec_deserialize(t *testing.T) {
	ser := Constructor()
	deser := Constructor()
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			data := ser.serialize(tt.args.root)
			if got := deser.deserialize(data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserializeP2(t *testing.T) {
	ser := Constructor()
	deser := Constructor()
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			data := ser.serializeP2(tt.args.root)
			if got := deser.deserializeP2(data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserializeP2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserializeP3(t *testing.T) {
	ser := Constructor()
	deser := Constructor()
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			data := ser.serializeP3(tt.args.root)
			if got := deser.deserializeP3(data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserializeP3() = %v, want %v", got, tt.want)
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
