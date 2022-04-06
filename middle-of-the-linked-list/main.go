/*
Package middle_of_the_linked_list
https://leetcode-cn.com/problems/middle-of-the-linked-list/

876. 链表的中间结点

给定一个头结点为 head 的非空单链表，返回链表的中间结点。
如果有两个中间结点，则返回第二个中间结点。

提示：
	给定链表的结点数介于 1 和 100 之间。

*/
package middle_of_the_linked_list

/**
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// --- 自己

/*
方法一: 快慢指针

时间复杂度：
空间复杂度：
*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
