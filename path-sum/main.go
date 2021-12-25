/*
Package path_sum
https://leetcode-cn.com/problems/path-sum/

112. 路径总和

给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
叶子节点 是指没有子节点的节点。

提示：
	树中节点的数目在范围 [0, 5000] 内
	-1000 <= Node.val <= 1000
	-1000 <= targetSum <= 1000

*/
package path_sum

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 深度优先搜索
	左节点和: targetSum - root.Val
	右节点和：targetSum - root.Val
	其中一个满足即满足

时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
空间复杂度：O(H)，其中 H 是树的高度。空间复杂度主要取决于递归时栈空间的开销，最坏情况下，树呈现链状，空间复杂度为 O(N)。
	平均情况下树的高度与节点数的对数正相关，空间复杂度为 O(logN)。
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	num := targetSum - root.Val
	if hasPathSum(root.Left, num) {
		return true
	}

	if hasPathSum(root.Right, num) {
		return true
	}

	return false
}
