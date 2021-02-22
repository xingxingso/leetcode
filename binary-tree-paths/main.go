/*
https://leetcode-cn.com/problems/binary-tree-paths/

257. 二叉树的所有路径

给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明:
	叶子节点是指没有子节点的节点。

*/
package binary_tree_paths

import "strconv"

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
func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	if root == nil {
		return ans
	}

	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			ans = append(ans, path)
			return
		}
		path += "->"
		dfs(root.Left, path)
		dfs(root.Right, path)
	}

	dfs(root, "")

	return ans
}
