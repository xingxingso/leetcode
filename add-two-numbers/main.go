/*
Package add_two_numbers
https://leetcode-cn.com/problems/add-two-numbers/

2. 两数相加

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

提示：
	每个链表中的节点数在范围 [1, 100] 内
	0 <= Node.val <= 9
	题目数据保证列表表示的数字不含前导零

*/
package add_two_numbers

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	add := 0
	node := ans
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + add
		node.Next = &ListNode{
			Val: sum % 10,
		}
		add = sum / 10
		l1 = l1.Next
		l2 = l2.Next
		node = node.Next
	}

	for l1 != nil {
		sum := l1.Val + add
		node.Next = &ListNode{
			Val: sum % 10,
		}
		add = sum / 10
		l1 = l1.Next
		node = node.Next
	}
	for l2 != nil {
		sum := l2.Val + add
		node.Next = &ListNode{
			Val: sum % 10,
		}
		add = sum / 10
		l2 = l2.Next
		node = node.Next
	}
	if add != 0 {
		node.Next = &ListNode{
			Val: add,
		}
	}

	return ans.Next
}
