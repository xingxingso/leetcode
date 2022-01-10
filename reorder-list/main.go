/*
Package reorder_list
https://leetcode-cn.com/problems/reorder-list/

143. 重排链表

给定一个单链表 L 的头节点 head ，单链表 L 表示为：
	L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：
	L0 → Ln → L1 → L_{n-1} → L2 → L_{n-2} → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

提示：
	链表的长度范围为 [1, 5 * 10^4]
	1 <= node.val <= 1000

*/
package reorder_list

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
	// 后半部分翻转
	// 依次插入前半部分

时间复杂度：
空间复杂度：
*/
func reorderList(head *ListNode) {
	// 找出后半部分
	slow, quick := head, head.Next // 获取前一个节点
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}

	//fmt.Println(slow.Val)

	// 翻转
	pre := slow
	for slow = slow.Next; slow != nil && slow.Next != nil; {
		node := slow.Next
		slow.Next = node.Next
		node.Next = pre.Next
		pre.Next = node
	}

	//printList(pre)

	//newHead := head
	//printList(newHead)

	// 插入
	for head != pre {
		// 断开 要插入的节点
		insert := pre.Next
		pre.Next = insert.Next

		cur := head
		head = head.Next
		cur.Next = insert
		insert.Next = head
		//fmt.Println("for:")
		//printList(newHead)
	}

	//printList(newHead)

	return
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println()
}
