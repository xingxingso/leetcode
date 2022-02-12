/*
Package subtree_of_another_tree
https://leetcode-cn.com/problems/subtree-of-another-tree/

572. 另一棵树的子树

给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。

二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。

提示：
	root 树上的节点数量范围是 [1, 2000]
	subRoot 树上的节点数量范围是 [1, 1000]
	-10^4 <= root.val <= 10^4
	-10^4 <= subRoot.val <= 10^4

*/
package subtree_of_another_tree

// TreeNode Definition for a binary tree node.
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(root, sub *TreeNode) bool
	dfs = func(root, sub *TreeNode) bool {
		if root == nil && sub == nil {
			return true
		}

		if root == nil || sub == nil {
			return false
		}

		return root.Val == sub.Val && dfs(root.Left, sub.Left) && dfs(root.Right, sub.Right)
	}

	if root == nil {
		return false
	}

	return dfs(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
