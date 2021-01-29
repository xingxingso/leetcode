package public

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
