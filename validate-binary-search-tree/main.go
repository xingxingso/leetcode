/*
Package validate_binary_search_tree
https://leetcode-cn.com/problems/validate-binary-search-tree/

98. 验证二叉搜索树

给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。

*/
package validate_binary_search_tree

import "math"

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一:
	利用中序遍历

时间复杂度：
空间复杂度：
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]int, 0)
	flag := true
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || !flag {
			return
		}
		inorder(root.Left)
		if len(stack) > 0 {
			last := stack[len(stack)-1]
			if root.Val <= last {
				flag = false
				return
			}
		}

		stack = append(stack, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return flag
}

// --- 官方

/*
方法一:
	利用中序遍历

时间复杂度：
空间复杂度：
*/
func isValidBSTO1(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
