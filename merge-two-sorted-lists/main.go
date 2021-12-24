/*
Package merge_two_sorted_lists
https://leetcode-cn.com/problems/merge-two-sorted-lists/

21. 合并两个有序链表

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

提示：
	两个链表的节点数目范围是 [0, 50]
	-100 <= Node.val <= 100
	l1 和 l2 均按 非递减顺序 排列

*/
package merge_two_sorted_lists

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for list1 != nil && list2 != nil {
		n := list1.Val
		if list1.Val > list2.Val {
			n = list2.Val
			list2 = list2.Next
		} else {
			list1 = list1.Next
		}
		head.Next = &ListNode{Val: n}
		head = head.Next
	}

	for ; list1 != nil; list1 = list1.Next {
		head.Next = &ListNode{Val: list1.Val}
		head = head.Next
	}
	for ; list2 != nil; list2 = list2.Next {
		head.Next = &ListNode{Val: list2.Val}
		head = head.Next
	}

	return dummy.Next
}
