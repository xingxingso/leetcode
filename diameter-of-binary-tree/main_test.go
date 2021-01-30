package diameter_of_binary_tree

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				root: root1(),
			},
			want: 3,
		},
		{
			name: "equal",
			args: args{
				root: root2(),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
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

func root2() *TreeNode {
	node5 := &TreeNode{}
	node5.Val = 5

	node3 := &TreeNode{}
	node3.Val = 3

	node2 := &TreeNode{}
	node2.Val = 2
	node2.Right = node5

	node1 := &TreeNode{}
	node1.Val = 1
	node1.Left = node2
	node1.Right = node3

	return node1
}
