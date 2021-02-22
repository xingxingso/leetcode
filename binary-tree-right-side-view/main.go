/*
https://leetcode-cn.com/problems/binary-tree-right-side-view/

199. 二叉树的右视图

给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

*/
package binary_tree_right_side_view

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
方法一: 广度优先搜索

时间复杂度：
空间复杂度：
*/
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ans = append(ans, queue[0].Val)
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
		}
		queue = queue[size:]
	}

	return ans
}

// --- 官方

/*
方法一: 深度优先搜索

时间复杂度：
空间复杂度：
*/
func rightSideViewS1(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(ans) == depth { // 关键部分
			ans = append(ans, root.Val)
		}
		depth++
		dfs(root.Right, depth)
		dfs(root.Left, depth) // 优先被 右子树 填充，所以当右子树节点未遍历完的时候，左子树不会被填充
	}

	dfs(root, 0)

	return ans
}
