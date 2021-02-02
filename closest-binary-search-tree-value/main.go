/*
https://leetcode-cn.com/problems/closest-binary-search-tree-value/

270. 最接近的二叉搜索树值

给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。

注意：
	给定的目标值 target 是一个浮点数
	题目保证在该二叉搜索树中只会存在一个最接近目标值的数

*/
package closest_binary_search_tree_value

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
方法一: 二分搜索

时间复杂度：O(H)
空间复杂度：O(1)
*/
func closestValue(root *TreeNode, target float64) int {
	minV := float64(root.Val) - target

	if abs(minV) <= 0.5 {
		return root.Val
	}

	if minV > 0.5 && root.Left != nil {
		mf := closestValue(root.Left, target)
		if abs(float64(mf)-target) < abs(minV) {
			return mf
		} else {
			return root.Val
		}
	}

	if minV < -0.5 && root.Right != nil {
		mr := closestValue(root.Right, target)
		if abs(float64(mr)-target) < abs(minV) {
			return mr
		} else {
			return root.Val
		}
	}

	return root.Val
}

func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
