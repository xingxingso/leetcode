/*
Package sort_list
https://leetcode-cn.com/problems/sort-list/

148. 排序链表

给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

提示：
	链表中节点的数目在范围 [0, 5 * 10^4] 内
	-10^5 <= Node.val <= 10^5

进阶：
	你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

*/
package sort_list

import "fmt"

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// --- 自己

/*
方法一: 快速排序
	// 超时时限

时间复杂度：
空间复杂度：
*/
func sortList(head *ListNode) *ListNode {
	var partition func(head, end *ListNode) (*ListNode, *ListNode)
	partition = func(head, end *ListNode) (*ListNode, *ListNode) {
		dummy := &ListNode{}
		h := dummy
		part, p := head, head
		for head = head.Next; head != end; head = head.Next {
			if head.Val < part.Val {
				h.Next = head
				h = h.Next
			} else {
				p.Next = head
				p = p.Next
			}
		}
		p.Next = end
		h.Next = part
		//printList(dummy.Next)
		return dummy.Next, part
	}

	var sort func(head, end *ListNode) *ListNode
	sort = func(head, end *ListNode) *ListNode {
		if head == nil || head.Next == nil || head == end {
			return head
		}
		h, p := partition(head, end)
		//fmt.Println("p, h", p.Val, h.Val)
		//printList(h)
		l1 := sort(h, p)
		//fmt.Println("l1")
		//printList(l1)
		l2 := sort(p.Next, nil)
		//fmt.Println("l2")
		//printList(l2)
		p.Next = l2
		return l1
	}
	first := sort(head, nil)
	//fmt.Println("result")
	//printList(first)
	return first
}

// --- 官方

/*
方法一: 自底向上归并排序

时间复杂度：O(n log n)，其中 n 是链表的长度。
空间复杂度：O(1)。
*/
func sortListO1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println()
}
