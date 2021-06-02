/*
Package remove_duplicates_from_sorted_list
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

83. 删除排序链表中的重复元素

存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
返回同样按升序排列的结果链表。

提示：
	链表中节点数目在范围 [0, 300] 内
	-100 <= Node.val <= 100
	题目数据保证链表已经按升序排列

*/
package remove_duplicates_from_sorted_list

import "fmt"

/*
ListNode
Definition for singly-linked list.
*/
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
func deleteDuplicates(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	if slow != nil {
		slow.Next = nil
	}
	return head
}

func deleteDuplicatesS1(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}
