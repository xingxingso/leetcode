/*
Package binary_tree_preorder_traversal
https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

144. 二叉树的前序遍历

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

提示：
	树中节点数目在范围 [0, 100] 内
	-100 <= Node.val <= 100

进阶：
	递归算法很简单，你可以通过迭代算法完成吗？

*/
package binary_tree_preorder_traversal

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 递归

时间复杂度：O(n)，其中 n 是二叉树的节点数。每一个节点恰好被遍历一次。
空间复杂度：O(n)，为递归过程中栈的开销，平均情况下为 O(logn)，最坏情况下树呈现链状，为 O(n)。
*/
func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, preorderTraversal(root.Left)...)
	ans = append(ans, preorderTraversal(root.Right)...)
	return ans
}

/*
方法二: 迭代
	借助栈，先压入右节点 后压入左节点

时间复杂度：
空间复杂度：
*/
func preorderTraversalS2(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}

// --- 官方

/*
方法二: 迭代

时间复杂度：O(n)，其中 n 是二叉树的节点数。每一个节点恰好被遍历一次。
空间复杂度：O(n)，为迭代过程中显式栈的开销，平均情况下为 O(logn)，最坏情况下树呈现链状，为 O(n)。
*/
func preorderTraversalO2(root *TreeNode) []int {
	vals := make([]int, 0)
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return vals
}

/*
方法三: Morris 遍历

时间复杂度：O(n)，其中 n 是二叉树的节点数。没有左子树的节点只被访问一次，有左子树的节点被访问两次。
空间复杂度：O(1)。只操作已经存在的指针（树的空闲指针），因此只需要常数的额外空间。
*/
func preorderTraversalO3(root *TreeNode) []int {
	vals := make([]int, 0)
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				vals = append(vals, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			vals = append(vals, p1.Val)
		}
		p1 = p1.Right
	}
	return vals
}
