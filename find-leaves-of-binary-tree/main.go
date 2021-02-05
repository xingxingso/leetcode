/*
https://leetcode-cn.com/problems/find-leaves-of-binary-tree/

366. 寻找二叉树的叶子节点

给你一棵二叉树，请按以下要求的顺序收集它的全部节点：
	1. 依次从左到右，每次收集并删除所有的叶子节点
	2. 重复如上过程直到整棵树为空

*/
package find_leaves_of_binary_tree

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func findLeaves(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	if isLeaves(root) {
		// 加入
	}

	if root.Left != nil {
		findLeaves(root.Left)
	}

	if
}

func isLeaves(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
