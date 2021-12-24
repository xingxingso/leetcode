/*
Package merge_two_binary_trees
https://leetcode-cn.com/problems/merge-two-binary-trees/

617. 合并二叉树

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
否则不为 NULL 的节点将直接作为新二叉树的节点。

注意:
	合并必须从两个树的根节点开始。

*/
package merge_two_binary_trees

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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	root := new(TreeNode)

	if root1 == nil && root2 != nil {
		root.Val = root2.Val
		root.Left = mergeTrees(nil, root2.Left)
		root.Right = mergeTrees(nil, root2.Right)
	} else if root1 != nil && root2 == nil {
		root.Val = root1.Val
		root.Left = mergeTrees(root1.Left, nil)
		root.Right = mergeTrees(root1.Right, nil)
	} else if root1 != nil && root2 != nil {
		root.Val = root1.Val + root2.Val
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	}

	return root
}

// --- 官方

/*
方法一: 深度优先搜索

时间复杂度：O(min(m,n))，其中 m 和 n 分别是两个二叉树的节点个数。对两个二叉树同时进行深度优先搜索，
	只有当两个二叉树中的对应节点都不为空时才会对该节点进行显性合并操作，因此被访问到的节点数不会超过较小的二叉树的节点数。
空间复杂度：O(min(m,n))，其中 m 和 n 分别是两个二叉树的节点个数。空间复杂度取决于递归调用的层数，
	递归调用的层数不会超过较小的二叉树的最大高度，最坏情况下，二叉树的高度等于节点数。
*/
func mergeTreesO1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
