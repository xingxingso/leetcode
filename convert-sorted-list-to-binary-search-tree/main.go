/*
Package convert_sorted_array_to_binary_search_tree
https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/

109. 有序链表转换二叉搜索树

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

*/
package convert_sorted_array_to_binary_search_tree

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一:
	快慢指针找到中间，中序遍历是顺序的

时间复杂度：
空间复杂度：
*/
func sortedListToBST(head *ListNode) *TreeNode {
	var dfs func(start, end *ListNode) *TreeNode
	dfs = func(start, end *ListNode) *TreeNode {
		if start == end {
			return nil
		}
		mid := getListMid(start, end)
		root := &TreeNode{
			Val: mid.Val,
		}
		root.Left = dfs(start, mid)
		root.Right = dfs(mid.Next, end)

		return root
	}

	return dfs(head, nil)
}

func getListMid(start, end *ListNode) *ListNode {
	if start == end {
		return start
	}
	slow, fast := start, start
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
