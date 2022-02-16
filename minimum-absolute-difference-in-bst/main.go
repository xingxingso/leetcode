/*
Package minimum_absolute_difference_in_bst
https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

530. 二叉搜索树的最小绝对差

给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。

提示：
	树中节点的数目范围是 [2, 10^4]
	0 <= Node.val <= 10^5

注意：
	本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同

*/
package minimum_absolute_difference_in_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一:
	利用中序遍历是有序数组

时间复杂度：
空间复杂度：
*/
func getMinimumDifference(root *TreeNode) int {
	num := -100000
	ans := 100000
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}
		ans = min(ans, root.Val-num)
		num = root.Val
		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
