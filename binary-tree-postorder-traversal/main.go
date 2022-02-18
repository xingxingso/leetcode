/*
Package binary_tree_postorder_traversal
https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

145. 二叉树的后序遍历

给定一个二叉树，返回它的 后序 遍历。

进阶:
	递归算法很简单，你可以通过迭代算法完成吗？

*/
package binary_tree_postorder_traversal

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 递归

时间复杂度：
空间复杂度：
*/
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

/*
方法二: 迭代

时间复杂度：
空间复杂度：
*/
func postorderTraversalS2(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append([]int{node.Val}, ans...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return ans
}

// --- 官方

/*
方法二: 迭代

时间复杂度：O(n)，其中 n 是二叉搜索树的节点数。每一个节点恰好被遍历一次。
空间复杂度：O(n)，为递归过程中栈的开销，平均情况下为 O(logn)，最坏情况下树呈现链状，为 O(n)。
*/
func postorderTraversalO2(root *TreeNode) []int {
	res := make([]int, 0)
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return res
}

/*
方法三: Morris 遍历

时间复杂度：O(n)，其中 n 是二叉树的节点数。没有左子树的节点只被访问一次，有左子树的节点被访问两次。
空间复杂度：O(1)。只操作已经存在的指针（树的空闲指针），因此只需要常数的额外空间。
*/
func postorderTraversalO3(root *TreeNode) (res []int) {
	addPath := func(node *TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
