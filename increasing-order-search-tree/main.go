/*
https://leetcode-cn.com/problems/increasing-order-search-tree/

897. 递增顺序查找树

给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。

提示：
	给定树中的结点数介于 1 和 100 之间。
	每个结点都有一个从 0 到 1000 范围内的唯一整数值。

*/
package increasing_order_search_tree

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
方法一: 递归
	复杂度分析来自官方, 官方解法使用了哨兵节点

时间复杂度：O(N)，其中 N 是树上的节点个数
空间复杂度：O(H)，其中 H 是数的高度。
*/
func increasingBST(root *TreeNode) *TreeNode {
	var head, cur *TreeNode
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		node.Left = nil
		if head == nil {
			head = node
			cur = head
		} else {
			cur.Right = node
			cur = cur.Right
		}
		inorder(node.Right)
	}

	inorder(root)
	return head
}

func Bfs(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, 0)
			continue
		}

		res = append(res, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	// 移除右边的0
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == 0 {
			continue
		}

		return res[:i+1]
	}

	return res
}
