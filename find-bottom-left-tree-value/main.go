/*
Package find_bottom_left_tree_value
https://leetcode-cn.com/problems/find-bottom-left-tree-value/

513. 找树左下角的值

给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。

提示:
	二叉树的节点个数的范围是 [1,10^4]
	-2^31 <= Node.val <= 2^31 - 1

*/
package find_bottom_left_tree_value

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 层次遍历

时间复杂度：
空间复杂度：
*/
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if i == 0 {
				ans = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}

	return ans
}
