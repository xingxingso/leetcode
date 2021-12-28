/*
Package binary_tree_inorder_traversal
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

94. 二叉树的中序遍历

给定一个二叉树的根节点 root ，返回它的 中序 遍历。

提示：
	树中节点数目在范围 [0, 100] 内
	-100 <= Node.val <= 100

进阶:
	递归算法很简单，你可以通过迭代算法完成吗？

*/
package binary_tree_inorder_traversal

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 递归

时间复杂度：O(n)，其中 n 为二叉树节点的个数。二叉树的遍历中每个节点会被访问一次且只会被访问一次。
空间复杂度：O(n)。空间复杂度取决于递归的栈深度，而栈深度在二叉树为一条链的情况下会达到 O(n) 的级别。
*/
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root != nil {
		ans = append(ans, inorderTraversal(root.Left)...)
		ans = append(ans, root.Val)
		ans = append(ans, inorderTraversal(root.Right)...)
	}
	return ans
}

/*
方法二: 迭代

时间复杂度：O(n)，其中 n 为二叉树节点的个数。二叉树的遍历中每个节点会被访问一次且只会被访问一次。
空间复杂度：O(n)。空间复杂度取决于栈深度，而栈深度在二叉树为一条链的情况下会达到 O(n) 的级别。
*/
func inorderTraversalS2(root *TreeNode) []int {
	var stack []*TreeNode
	ans := make([]int, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		root = node.Right
	}
	return ans
}

// --- 官方

/*
方法三: Morris 中序遍历

时间复杂度：O(n)，其中 n 为二叉搜索树的节点个数。Morris 遍历中每个节点会被访问两次，因此总时间复杂度为 O(2n)=O(n)。
空间复杂度：O(1)。
*/
func inorderTraversalO3(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return res
}
