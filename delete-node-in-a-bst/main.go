/*
Package delete_node_in_a_bst
https://leetcode-cn.com/problems/delete-node-in-a-bst/

450. 删除二叉搜索树中的节点

给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
	1. 首先找到需要删除的节点；
	2. 如果找到了，删除它。
说明：
	要求算法时间复杂度为 O(h)，h 为树的高度。

*/
package delete_node_in_a_bst

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
方法一: 深度优先搜索

时间复杂度：
空间复杂度：
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left != nil && root.Right != nil {
			root.Right = insertNode(root.Right, root.Left)
			return root.Right
		} else if root.Left != nil {
			return root.Left
		} else { // left right 同时为 nil 同样可以返回 root.Right
			return root.Right
		}
	}
	return root
}

func insertNode(root, in *TreeNode) *TreeNode {
	var dfs func(root, in *TreeNode) *TreeNode
	dfs = func(root, in *TreeNode) *TreeNode {
		if root == nil {
			return in
		}
		if root.Val > in.Val {
			root.Left = dfs(root.Left, in)
		} else {
			root.Right = dfs(root.Right, in)
		}
		return root
	}
	return dfs(root, in)
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247488128&idx=2&sn=b8fb3fd2917f9ac86127054741cd5877&chksm=9bd7ec88aca0659ee0185b657663169169493e9df2063fa4d28b38a0b4d0dd698d0301937898&scene=21#wechat_redirect

/*
方法一:

时间复杂度：
空间复杂度：
*/
func deleteNodeP1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNodeP1(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNodeP1(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNodeP1(root.Right, minNode.Val)
	}
	return root
}

// BST 最左边的就是最小的
func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
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
