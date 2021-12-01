/*
Package kth_smallest_element_in_a_bst
https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

230. 二叉搜索树中第K小的元素

给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

提示：
	树中的节点数为 n 。
	1 <= k <= n <= 10^4
	0 <= Node.val <= 10^4

进阶：
	如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

*/
package kth_smallest_element_in_a_bst

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
方法一: 深度优先搜索

时间复杂度：
空间复杂度：
*/
func kthSmallest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	res := -1
	dfs = func(root *TreeNode) {
		if root.Left != nil && res == -1 {
			dfs(root.Left)
		}
		k--
		if k == 0 {
			res = root.Val
		}
		if root.Right != nil && res == -1 {
			dfs(root.Right)
		}
		return
	}
	dfs(root)
	return res
}

// --- 官方

/*
方法二: 迭代

时间复杂度：O(H+k)，其中 HH 指的是树的高度，由于我们开始遍历之前，要先向下达到叶，当树是一个平衡树时：
	复杂度为 O(logN+k)。当树是一个不平衡树时：复杂度为 O(N+k)，此时所有的节点都在左子树。
空间复杂度：O(H+k)。当树是一个平衡树时：O(logN+k)。当树是一个非平衡树时：O(N+k)。
*/
func kthSmallestO1(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
