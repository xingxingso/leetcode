/*
Package sum_of_left_leaves
https://leetcode-cn.com/problems/sum-of-left-leaves/

404. 左叶子之和

给定二叉树的根节点 root ，返回所有左叶子之和。

提示:
	节点数在 [1, 1000] 范围内
	-1000 <= Node.val <= 1000

*/
package sum_of_left_leaves

// TreeNode Definition for a binary tree node.
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
func sumOfLeftLeaves(root *TreeNode) int {
	ans := 0
	//var dfs func(root *TreeNode)
	//dfs = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//	if root.Left != nil {
	//		if root.Left.Left == nil && root.Left.Right == nil {
	//			ans += root.Left.Val
	//		} else {
	//			dfs(root.Left)
	//		}
	//	}
	//
	//	if root.Right != nil {
	//		dfs(root.Right)
	//	}
	//}
	//
	//dfs(root)

	if root == nil {
		return 0
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			ans += root.Left.Val
		} else {
			ans += sumOfLeftLeaves(root.Left)
		}
	}

	if root.Right != nil {
		ans += sumOfLeftLeaves(root.Right)
	}

	return ans
}
