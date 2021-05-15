/*
https://leetcode-cn.com/problems/linked-list-cycle-ii/

142. 环形链表 II

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：
	不允许修改给定的链表。

进阶：
	你是否可以使用 O(1) 空间解决此题？

提示：
	链表中节点的数目范围在范围 [0, 10^4] 内
	-10^5 <= Node.val <= 10^5
	pos 的值为 -1 或者链表中的一个有效索引

*/
package linked_list_cycle_ii

/**
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
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	slow = head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return fast
		}
		slow = slow.Next
		fast = fast.Next
	}
	return nil
}
