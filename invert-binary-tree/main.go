/*
https://leetcode-cn.com/problems/invert-binary-tree/

226. 翻转二叉树

翻转一棵二叉树。

备注:
	这个问题是受到 Max Howell 的 原问题 启发的 ：
	谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
*/
package invert_binary_tree

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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
