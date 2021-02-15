/*
https://leetcode-cn.com/problems/find-all-the-lonely-nodes/

1469. 寻找所有的独生节点

二叉树中，如果一个节点是其父节点的唯一子节点，则称这样的节点为 “独生节点” 。二叉树的根节点不会是独生节点，因为它没有父节点。
给定一棵二叉树的根节点 root ，返回树中 所有的独生节点的值所构成的数组 。数组的顺序 不限 。

提示：
	tree 中节点个数的取值范围是 [1, 1000]。
	每个节点的值的取值范围是 [1, 10^6]。

*/
package find_all_the_lonely_nodes

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
func getLonelyNodes(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	if root.Left == nil && root.Right != nil {
		ans = append(ans, root.Right.Val)
	} else if root.Left != nil && root.Right == nil {
		ans = append(ans, root.Left.Val)
	}

	if root.Left != nil {
		left := getLonelyNodes(root.Left)
		ans = append(ans, left...)
	}

	if root.Right != nil {
		right := getLonelyNodes(root.Right)
		ans = append(ans, right...)
	}

	return ans
}
