/*
Package intersection_of_two_linked_lists
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

160. 相交链表

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

自定义评测：

评测系统 的输入如下（你设计的程序 不适用 此输入）：
	intersectVal - 相交的起始节点的值。如果不存在相交节点，这一值为 0
	listA - 第一个链表
	listB - 第二个链表
	skipA - 在 listA 中（从头节点开始）跳到交叉节点的节点数
	skipB - 在 listB 中（从头节点开始）跳到交叉节点的节点数
	评测系统将根据这些输入创建链式数据结构，并将两个头节点 headA 和 headB 传递给你的程序。如果程序能够正确返回相交节点，那么你的解决方案将被 视作正确答案 。

提示：
	listA 中节点数目为 m
	listB 中节点数目为 n
	1 <= m, n <= 3 * 10^4
	1 <= Node.val <= 10^5
	0 <= skipA <= m
	0 <= skipB <= n
	如果 listA 和 listB 没有交点，intersectVal 为 0
	如果 listA 和 listB 有交点，intersectVal == listA[skipA] == listB[skipB]

进阶：
	你能否设计一个时间复杂度 O(m + n) 、仅用 O(1) 内存的解决方案？

*/
package intersection_of_two_linked_lists

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m, n := 0, 0
	for head := headA; head != nil; head = head.Next {
		m++
	}
	for head := headB; head != nil; head = head.Next {
		n++
	}

	dif := 0
	var start, follow *ListNode
	if m > n {
		dif = m - n
		start = headA
		follow = headB
	} else {
		dif = n - m
		start = headB
		follow = headA
	}
	for ; dif > 0; dif-- {
		start = start.Next
	}

	for start != nil && start != follow {
		start = start.Next
		follow = follow.Next
	}

	if start == nil && follow == nil {
		return nil
	}

	return start
}

// --- 官方

/*
方法一: 双指针

时间复杂度：
空间复杂度：
*/
func getIntersectionNodeO1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
