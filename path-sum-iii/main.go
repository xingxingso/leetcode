/*
Package path_sum_iii
https://leetcode-cn.com/problems/path-sum-iii/

437. 路径总和 III

给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

提示:
	二叉树的节点个数的范围是 [0,1000]
	-10^9 <= Node.val <= 10^9
	-1000 <= targetSum <= 1000

*/
package path_sum_iii

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
func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}
		if root.Left == nil && root.Right == nil {
			if root.Val == targetSum {
				ans++
			}
			return []int{root.Val}
		}
		res := []int{root.Val}
		if root.Val == targetSum {
			ans++
		}
		if root.Left != nil {
			for _, v := range dfs(root.Left) {
				k := v + root.Val
				if k == targetSum {
					ans++
				}
				res = append(res, k)
			}
		}
		if root.Right != nil {
			for _, v := range dfs(root.Right) {
				k := v + root.Val
				if k == targetSum {
					ans++
				}
				res = append(res, k)
			}
		}
		return res
	}

	_ = dfs(root)

	return ans
}

// --- 官方

/*
方法一: 深度优先搜索

时间复杂度：
空间复杂度：
*/
func pathSumO1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var dfs func(root *TreeNode, targetSum int) int
	dfs = func(root *TreeNode, targetSum int) int {
		res := 0
		if root == nil {
			return res
		}
		if root.Val == targetSum {
			res++
		}
		res += dfs(root.Left, targetSum-root.Val)
		res += dfs(root.Right, targetSum-root.Val)
		return res
	}
	res := dfs(root, targetSum)
	res += pathSumO1(root.Left, targetSum)
	res += pathSumO1(root.Right, targetSum)
	return res
}
