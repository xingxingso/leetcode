/*
Package maximum_depth_of_binary_tree
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

104. 二叉树的最大深度

给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明:
	叶子节点是指没有子节点的节点。

*/
package maximum_depth_of_binary_tree

//TreeNode Definition for a binary tree node.
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
func maxDepth(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			ans = max(ans, depth)
			return
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// --- 官方

/*
方法一: 深度优先搜索

时间复杂度：O(n)，其中 n 为二叉树节点的个数。每个节点在递归中只被遍历一次。
空间复杂度：O(height)，其中 height 表示二叉树的高度。递归函数需要栈空间，而栈空间取决于递归的深度，因此空间复杂度等价于二叉树的高度。
*/
func maxDepthO1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthO1(root.Left), maxDepthO1(root.Right)) + 1
}
