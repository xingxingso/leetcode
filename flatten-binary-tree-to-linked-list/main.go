/*
https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

114. 二叉树展开为链表

给你二叉树的根结点 root ，请你将它展开为一个单链表：
	- 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	- 展开后的单链表应该与二叉树 先序遍历 顺序相同。

提示：
	树中结点数在范围 [0, 2000] 内
	-100 <= Node.val <= 100

进阶：
	你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

*/
package flatten_binary_tree_to_linked_list

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

时间复杂度：
空间复杂度：
*/
func flatten(root *TreeNode) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || root.Left == nil && root.Right == nil {
			return
		}
		head := root
		dfs(head.Left)
		dfs(head.Right)
		right := head.Right
		left := head.Left
		head.Left = nil
		head.Right = left
		for head.Right != nil {
			head = head.Right
		}
		head.Right = right
	}
	dfs(root)
}

// --- 官方

/*
方法三：寻找前驱节点

时间复杂度：
空间复杂度：
*/
func flattenO1(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
