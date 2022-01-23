/*
Package range_sum_of_bst
https://leetcode-cn.com/problems/range-sum-of-bst/

938. 二叉搜索树的范围和

给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

提示：
	树中节点数目在范围 [1, 2 * 104] 内
	1 <= Node.val <= 105
	1 <= low <= high <= 105
	所有 Node.val 互不相同

*/
package range_sum_of_bst

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
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	ans := 0

	if root.Val >= low && root.Val <= high {
		ans += root.Val
	}

	if root.Val > low && root.Left != nil {
		ans += rangeSumBST(root.Left, low, high)
	}
	if root.Val < high && root.Right != nil {
		ans += rangeSumBST(root.Right, low, high)
	}

	return ans
}
