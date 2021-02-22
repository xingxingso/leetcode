/*
https://leetcode-cn.com/problems/path-sum-ii/

113. 路径总和 II

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明:
	叶子节点是指没有子节点的节点。

*/
package path_sum_ii

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
方法一: 深度优先遍历

时间复杂度：
空间复杂度：
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	var dfs func(root *TreeNode, path []int, targetSum int)
	dfs = func(root *TreeNode, path []int, targetSum int) {
		if root == nil {
			return
		}
		newTarget := targetSum - root.Val
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil && newTarget == 0 {
			ans = append(ans, path)
			return
		}

		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		dfs(root.Left, path, newTarget)
		dfs(root.Right, pathCopy, newTarget)
		// dfs(root.Right, path, newTarget) // 不能直接这么写
	}
	dfs(root, []int{}, targetSum)
	return ans
}
