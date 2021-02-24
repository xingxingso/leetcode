/*
https://leetcode-cn.com/problems/diameter-of-binary-tree/

543. 二叉树的直径

给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

注意：
	两结点之间的路径长度是以它们之间边的数目表示。
*/
package diameter_of_binary_tree

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
方法一：

时间复杂度：
空间复杂度：
*/
func diameterOfBinaryTreeS1(root *TreeNode) int {
	ans := 0
	var depth func(root *TreeNode, dep int) int
	depth = func(root *TreeNode, dep int) int {
		if root == nil {
			return 0
		}
		left := depth(root.Left, dep)
		right := depth(root.Right, dep)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}
	depth(root, 0)
	return ans
}

// --- 官方

/*
方法一：深度优先搜索
	一条路径的长度为该路径经过的节点数减一，所以求直径（即求路径长度的最大值）等效于求路径经过节点数的最大值减一。

时间复杂度：O(N)，其中 N 为二叉树的节点数，即遍历一棵二叉树的时间复杂度，每个结点只被访问一次。
空间复杂度：O(Height)，其中 Height 为二叉树的高度。由于递归函数在递归过程中需要为每一层递归函数分配栈空间，
	所以这里需要额外的空间且该空间取决于递归的深度，而递归的深度显然为二叉树的高度，并且每次递归调用的函数里又只用了常数个变量，
	所以所需空间复杂度为 O(Height) 。
*/
func diameterOfBinaryTree(root *TreeNode) int {
	var depth func(node *TreeNode) int
	ans := 1
	depth = func(node *TreeNode) int {
		// 访问到空节点了，返回0
		if node == nil {
			return 0
		}
		// 左儿子为根的子树的深度
		l := depth(node.Left)
		// 右儿子为根的子树的深度
		r := depth(node.Right)
		// 计算d_node即L+R+1 并更新ans
		ans = max(ans, l+r+1)
		// 返回该节点为根的子树的深度
		return max(l, r) + 1
	}
	depth(root)
	return ans - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
