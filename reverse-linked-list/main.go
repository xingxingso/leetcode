/*
https://leetcode-cn.com/problems/reverse-linked-list/

206. 反转链表

反转一个单链表。

进阶：
	你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

*/
package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// --- 自己

/*
方法一: 迭代

时间复杂度：O(n)，假设 n 是列表的长度，时间复杂度是 O(n)。
空间复杂度：O(1)。
*/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		node := cur.Next
		cur.Next = pre
		pre = cur
		cur = node
	}

	return pre
}

// --- 官方

/*
方法二: 递归

时间复杂度：O(n)，假设 n 是列表的长度，那么时间复杂度为 O(n)。
空间复杂度：O(n)，由于使用递归，将会使用隐式栈空间。递归深度可能会达到 n 层。
*/
func reverseListO1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListO1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
