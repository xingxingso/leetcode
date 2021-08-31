/*
Package convert_bst_to_greater_tree
https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

538. 把二叉搜索树转换为累加树

给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
提醒一下，二叉搜索树满足下列约束条件：

	节点的左子树仅包含键 小于 节点键的节点。
	节点的右子树仅包含键 大于 节点键的节点。
	左右子树也必须是二叉搜索树。
注意：
	本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同

提示：
	树中的节点数介于 0 和 10^4 之间。
	每个节点的值介于 -10^4 和 10^4 之间。
	树中的所有值 互不相同 。
	给定的树为二叉搜索树。

*/
package convert_bst_to_greater_tree

import "fmt"

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

时间复杂度：
空间复杂度：
*/
func convertBST(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	sum := 0
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
		return
	}
	// inOrderPrint(root)
	// fmt.Println()
	dfs(root)
	// inOrderPrint(root)
	// fmt.Println()
	return root
}

// --- 官方

/*
方法二: Morris 遍历

时间复杂度：O(n)，其中 n 是二叉搜索树的节点数。没有左子树的节点只被访问一次，有左子树的节点被访问两次。
空间复杂度：O(1)。只操作已经存在的指针（树的空闲指针），因此只需要常数的额外空间。
*/
func convertBSTO1(root *TreeNode) *TreeNode {
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}

func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

func inOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderPrint(root.Left)
	fmt.Printf("%d,", root.Val)
	inOrderPrint(root.Right)
	// fmt.Println()
}
