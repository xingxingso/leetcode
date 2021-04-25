/*
https://leetcode-cn.com/problems/reverse-linked-list-ii/

92. 反转链表 II

给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

提示：
	链表中节点数目为 n
	1 <= n <= 500
	-500 <= Node.val <= 500
	1 <= left <= right <= n

进阶：
	你可以使用一趟扫描完成反转吗？

*/
package reverse_linked_list_ii

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// --- 自己

/*
方法一:
	参考 labuladong

时间复杂度：
空间复杂度：
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	for i := 1; head != nil && i < left; i++ {
		pre = pre.Next
	}

	node := reverseN(pre.Next, right-left+1)

	if left == 1 {
		return node
	}
	pre.Next = node
	return head
}

func reverseBetweenS2(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		return reverseN(head, right)
	}
	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 反转以 head 为起点的 n 个节点，返回新的头结点
func reverseN(head *ListNode, n int) *ListNode {
	var reverse func(head *ListNode, n int) *ListNode
	// 后驱节点
	var nextNode *ListNode
	reverse = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			// 记录第 n + 1 个节点
			nextNode = head.Next
			return head
		}
		// 以 head.next 为起点，需要反转前 n - 1 个节点
		last := reverse(head.Next, n-1)
		head.Next.Next = head
		// 让反转之后的 head 节点和后面的节点连起来
		head.Next = nextNode
		return last
	}

	return reverse(head, n)
}

func reverseBetweenS3(head *ListNode, left int, right int) *ListNode {
	var reverse func(head *ListNode, left, right int) *ListNode
	var nextNode *ListNode
	reverse = func(head *ListNode, left, right int) *ListNode {
		if left > 1 {
			node := reverse(head.Next, left-1, right-1)
			head.Next = node
			return head
		}

		if right == 1 {
			nextNode = head.Next
			return head
		}

		last := reverse(head.Next, left-1, right-1)
		head.Next.Next = head
		head.Next = nextNode
		return last
	}

	return reverse(head, left, right)
}

// --- 官方

/*
方法二: 一次遍历「穿针引线」反转链表（头插法）

时间复杂度：O(N)，其中 N 是链表总节点数。最多只遍历了链表一次，就完成了反转。
空间复杂度：O(1)。只使用到常数个变量。
*/
func reverseBetweenO1(head *ListNode, left int, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}
