package binary_tree_upside_down

import (
	"reflect"
	"testing"
)

func Test_upsideDownBinaryTreeP1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "equal",
			args: args{
				root: root1(),
			},
			want: want1(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upsideDownBinaryTreeP1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("upsideDownBinaryTreeP1() = %v, want %v", got, tt.want)
			}
			// 放在一起上一次执行已变更
			// if got := upsideDownBinaryTreeP2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("upsideDownBinaryTreeP2() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_upsideDownBinaryTreeP2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "equal",
			args: args{
				root: root1(),
			},
			want: want1(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upsideDownBinaryTreeP2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("upsideDownBinaryTreeP2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func root1() *TreeNode {
	node5 := &TreeNode{}
	node5.Val = 5

	node4 := &TreeNode{}
	node4.Val = 4

	node3 := &TreeNode{}
	node3.Val = 3

	node2 := &TreeNode{}
	node2.Val = 2
	node2.Left = node4
	node2.Right = node5

	node1 := &TreeNode{}
	node1.Val = 1
	node1.Left = node2
	node1.Right = node3

	return node1
}

func want1() *TreeNode {
	node3 := &TreeNode{}
	node3.Val = 3

	node1 := &TreeNode{}
	node1.Val = 1

	node5 := &TreeNode{}
	node5.Val = 5

	node2 := &TreeNode{}
	node2.Val = 2
	node2.Left = node3
	node2.Right = node1

	node4 := &TreeNode{}
	node4.Val = 4
	node4.Left = node5
	node4.Right = node2

	return node4
}
