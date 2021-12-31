/*
Package reverse_nodes_in_k_group
https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

25. K 个一组翻转链表

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

说明：
	你的算法只能使用常数的额外空间。
	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

*/
package reverse_nodes_in_k_group

import "fmt"

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
func reverseKGroupS1(head *ListNode, k int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	for pre := dummyNode; head != nil && head.Next != nil; head = head.Next {
		// 计数器 记录每个节点的序号
		count := 1
		// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
		for temp := head; count < k && temp.Next != nil; count++ {
			temp = temp.Next
		}

		if count < k {
			break
		}

		for count = 1; count < k && head.Next != nil; count++ {
			next := head.Next
			head.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}

		pre = head
	}

	return dummyNode.Next
}

// --- 官方

/*
方法一: 模拟

时间复杂度：O(n)，其中 n 为链表的长度。head 指针会在 O({n}/{k}) 个结点上停留，每次停留需要进行一次 O(k) 的翻转操作。
空间复杂度：O(1)，我们只需要建立常数个变量。
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}

	return tail, head
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484597&idx=1&sn=c603f1752e33cb2701e371d84254aee2&chksm=9bd7fabdaca073abd512d8fff18016c9092ede45fed65c307852c65a2026d8568ee294563c78&scene=21#wechat_redirect

/*
方法一:

时间复杂度：
空间复杂度：
*/
func reverseKGroupP1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反转元素
	a, b := head, head
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前 k 个元素
	newHead := reverse(a, b)
	// 递归反转后续链表并连接起来
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur, nxt := a, a
	for cur != b {
		nxt = cur.Next
		// 逐个结点反转
		cur.Next = pre
		// 更新指针位置
		pre = cur
		cur = nxt
	}
	// 返回反转后的头结点
	return pre
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}
