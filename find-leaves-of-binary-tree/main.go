/*
https://leetcode-cn.com/problems/find-leaves-of-binary-tree/

366. 寻找二叉树的叶子节点

给你一棵二叉树，请按以下要求的顺序收集它的全部节点：
	1. 依次从左到右，每次收集并删除所有的叶子节点
	2. 重复如上过程直到整棵树为空

*/
package find_leaves_of_binary_tree

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 他人

/*
https://leetcode-cn.com/problems/find-leaves-of-binary-tree/solution/c-hou-xu-bian-li-fan-xiang-ding-yi-shen-du-by-fatk/
方法一: 后序遍历，反向定义深度

时间复杂度：O(n)
空间复杂度：
*/
func findLeavesP1(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		// 左
		left := dfs(node.Left)
		// 右
		right := dfs(node.Right)
		// 本节点
		cur := max(left, right) + 1
		if cur >= len(res) {
			res = append(res, []int{})
		}
		res[cur] = append(res[cur], node.Val)
		return cur
	}

	dfs(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
https://leetcode-cn.com/problems/find-leaves-of-binary-tree/solution/di-gui-100-by-gennji61/
方法一: 递归

时间复杂度：
空间复杂度：
*/
func findLeavesP2(root *TreeNode) [][]int {
	var res [][]int
	var recur func(node *TreeNode, list *[]int) *TreeNode
	recur = func(node *TreeNode, list *[]int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			*list = append(*list, node.Val)
			return nil
		}
		node.Left = recur(node.Left, list)
		node.Right = recur(node.Right, list)
		return node
	}

	for root != nil {
		list := make([]int, 0)
		root = recur(root, &list)
		res = append(res, list)
	}

	return res
}
