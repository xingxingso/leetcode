/*
Package binary_tree_zigzag_level_order_traversal
https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

103. 二叉树的锯齿形层序遍历

给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

*/
package binary_tree_zigzag_level_order_traversal

// TreeNode Definition for a binary tree node.
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	flag := true
	for len(queue) > 0 {
		temp := make([]int, 0)
		l := len(queue)
		if flag {
			for i := 0; i < l; i++ {
				temp = append(temp, queue[i].Val)
			}
		} else {
			for i := l - 1; i >= 0; i-- {
				temp = append(temp, queue[i].Val)
			}
		}
		ans = append(ans, temp)

		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		flag = !flag
		queue = queue[l:]
	}

	return ans
}
