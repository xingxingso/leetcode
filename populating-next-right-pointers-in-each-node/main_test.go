package populating_next_right_pointers_in_each_node

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "equal0",
			args: args{
				root: CreateNodeByArray([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			want: CreateNode2(),
		},
		{
			name: "equal1",
			args: args{
				root: CreateNodeByArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			},
			want: CreateNode3(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}

			if got := connectS2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}

			if got := connectO1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func CreateNodeByArray(nums []int) *Node {
	if len(nums) <= 0 {
		return nil
	}
	queue := make([]*Node, 0)
	head := &Node{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums); {
		if nums[i] != 0 {
			queue[0].Left = &Node{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if i+1 < len(nums) && nums[i+1] != 0 {
			queue[0].Right = &Node{Val: nums[i+1]}
			queue = append(queue, queue[0].Right)
		}

		queue = queue[1:]
		i += 2
	}
	return head
}

func CreateNode2() *Node {
	node7 := &Node{
		Val: 7,
	}
	node6 := &Node{
		Val:  6,
		Next: node7,
	}
	node5 := &Node{
		Val:  5,
		Next: node6,
	}
	node4 := &Node{
		Val:  4,
		Next: node5,
	}
	node3 := &Node{
		Val:   3,
		Left:  node6,
		Right: node7,
	}
	node2 := &Node{
		Val:   2,
		Left:  node4,
		Right: node5,
		Next:  node3,
	}
	node1 := &Node{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	return node1
}

func CreateNode3() *Node {
	node15 := &Node{
		Val: 15,
	}
	node14 := &Node{
		Val:  14,
		Next: node15,
	}
	node13 := &Node{
		Val:  13,
		Next: node14,
	}
	node12 := &Node{
		Val:  12,
		Next: node13,
	}
	node11 := &Node{
		Val:  11,
		Next: node12,
	}
	node10 := &Node{
		Val:  10,
		Next: node11,
	}
	node9 := &Node{
		Val:  9,
		Next: node10,
	}
	node8 := &Node{
		Val:  8,
		Next: node9,
	}
	node7 := &Node{
		Val:   7,
		Left:  node14,
		Right: node15,
	}
	node6 := &Node{
		Val:   6,
		Left:  node12,
		Right: node13,
		Next:  node7,
	}
	node5 := &Node{
		Val:   5,
		Left:  node10,
		Right: node11,
		Next:  node6,
	}
	node4 := &Node{
		Val:   4,
		Left:  node8,
		Right: node9,
		Next:  node5,
	}
	node3 := &Node{
		Val:   3,
		Left:  node6,
		Right: node7,
	}
	node2 := &Node{
		Val:   2,
		Left:  node4,
		Right: node5,
		Next:  node3,
	}
	node1 := &Node{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	return node1
}
