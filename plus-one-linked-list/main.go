/*
https://leetcode-cn.com/problems/plus-one-linked-list/

369. 给单链表加一

用一个 非空 单链表来表示一个非负整数，然后将这个整数加一。
你可以假设这个整数除了 0 本身，没有任何前导的 0。
这个整数的各个数位按照 高位在链表头部、低位在链表尾部 的顺序排列。

*/
package plus_one_linked_list

/**
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// --- 官方

/*
方法一: 哨兵头节点

时间复杂度：O(N)，最多只需遍历两遍链表。
空间复杂度：O(1)。
*/
func plusOne(head *ListNode) *ListNode {
	// sentinel head
	sentinel := &ListNode{
		Val:  0,
		Next: head,
	}

	// find the rightmost not-nine digit
	notNine := sentinel
	for head != nil {
		if head.Val != 9 {
			notNine = head
		}
		head = head.Next
	}

	// increase this rightmost not-nine digit by 1
	notNine.Val++

	// set all the following nines to zeros
	for notNine.Next != nil {
		notNine = notNine.Next
		notNine.Val = 0
	}

	if sentinel.Val == 0 {
		return sentinel.Next
	}

	return sentinel
}
