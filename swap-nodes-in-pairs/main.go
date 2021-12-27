/*
Package swap_nodes_in_pairs
https://leetcode-cn.com/problems/swap-nodes-in-pairs/

24. 两两交换链表中的节点

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

提示：
	链表中节点的数目在范围 [0, 100] 内
	0 <= Node.val <= 100

*/
package swap_nodes_in_pairs

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := swapPairs(head.Next.Next)
	newHead := head.Next
	head.Next = tail
	newHead.Next = head

	// printList(newHead)
	return newHead
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}
