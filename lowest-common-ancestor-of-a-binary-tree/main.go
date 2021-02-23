/*
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

236. 二叉树的最近公共祖先

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

提示：
	树中节点数目在范围 [2, 105] 内。
	-109 <= Node.val <= 109
	所有 Node.val 互不相同 。
	p != q
	p 和 q 均存在于给定的二叉树中。

*/
package lowest_common_ancestor_of_a_binary_tree

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
	重复搜索了

时间复杂度：
空间复杂度：
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := lowestCommonAncestor(root.Left, p, q)
	if left != nil {
		return left
	}

	right := lowestCommonAncestor(root.Right, p, q)
	if right != nil {
		return right
	}

	stack := []*TreeNode{root}
	find := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node == p || node == q {
			find++
		}
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	if find == 2 {
		return root
	}

	return nil
}

// --- 官方

/*
方法一: 递归

时间复杂度：O(N)，其中 N 是二叉树的节点数。二叉树的所有节点有且只会被访问一次，因此时间复杂度为 O(N)。
空间复杂度：O(N) ，其中 N 是二叉树的节点数。递归调用的栈深度取决于二叉树的高度，二叉树最坏情况下为一条链，此时高度为 N，因此空间复杂度为 O(N)。
*/
func lowestCommonAncestorO1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestorO1(root.Left, p, q)
	right := lowestCommonAncestorO1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}
