/*
Package symmetric_tree
https://leetcode-cn.com/problems/symmetric-tree/

101. 对称二叉树

给定一个二叉树，检查它是否是镜像对称的

进阶：
	你可以运用递归和迭代两种方法解决这个问题吗？

*/
package symmetric_tree

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
	广度优先搜索， 搜索一层放入队列（按左右对称放入）， 取队列元素比较

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		num := len(queue)
		for i := 0; i < num; i += 2 {
			if queue[i] == nil && queue[i+1] == nil {
				continue
			}

			if queue[i] == nil && queue[i+1] != nil ||
				queue[i] != nil && queue[i+1] == nil {
				return false
			}

			if queue[i].Val != queue[i+1].Val {
				return false
			}

			queue = append(queue, queue[i].Left, queue[i+1].Right, queue[i].Right, queue[i+1].Left)
		}

		queue = queue[num:]
	}

	return true
}
