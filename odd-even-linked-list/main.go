/*
Package odd_even_linked_list
https://leetcode-cn.com/problems/odd-even-linked-list/

328. 奇偶链表

给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

提示:
	n ==  链表中的节点数
	0 <= n <= 10^4
	-10^6 <= Node.val <= 10^6

*/
package odd_even_linked_list

import "fmt"

//ListNode Definition for singly-linked list.
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
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even, evenHead := head, head.Next, head.Next
	for even != nil {
		odd.Next = even.Next
		if even.Next != nil {
			even.Next = even.Next.Next
		}
		even = even.Next
		if odd.Next != nil {
			odd = odd.Next
		}
	}
	//if odd != nil {
	odd.Next = evenHead
	//}
	//printList(head)
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println()
}
