/*
https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

124. 二叉树中的最大路径和

路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。
该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

提示：
	树中节点数目范围是 [1, 3 * 104]
	-1000 <= Node.val <= 1000

*/
package binary_tree_maximum_path_sum

import "math"

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
	计算左侧路径最大值 右侧路径最大值，
	max 实时更新为 left, left+parent, left+parent+right, parent, parent+right, right 中最大值
	自底向上

时间复杂度：O(N)，其中 N 是二叉树中的节点个数。对每个节点访问不超过 2 次。
空间复杂度：O(N)，其中 N 是二叉树中的节点个数。空间复杂度主要取决于递归调用层数，最大层数等于二叉树的高度，
	最坏情况下，二叉树的高度等于二叉树中的节点个数。
*/
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left, right := math.MinInt32, math.MinInt32
		if node.Left != nil {
			left = dfs(node.Left)
		}
		if node.Right != nil {
			right = dfs(node.Right)
		}
		ans = maxN(ans, left, left+node.Val, left+node.Val+right, node.Val, node.Val+right, right)
		return maxN(left+node.Val, node.Val, right+node.Val)
	}
	dfs(root)
	return ans
}

func maxN(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max
}
