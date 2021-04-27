/*
https://leetcode-cn.com/problems/palindrome-linked-list/

234. 回文链表

请判断一个链表是否为回文链表。

进阶：
	你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

*/
package palindrome_linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// --- 自己

/*
方法一:
	// 快慢指针找到中间位置， 翻转后半部分链表

时间复杂度：
空间复杂度：
*/
func isPalindrome(head *ListNode) bool {
	mid, tail := head, head

	// 奇数 mid 为中间值， 偶数为 中间靠前一个
	for tail.Next != nil && tail.Next.Next != nil {
		mid = mid.Next
		tail = tail.Next.Next
	}

	if mid == nil {
		return true
	}
	cur := mid.Next
	for cur != nil && cur.Next != nil {
		nex := cur.Next
		cur.Next = nex.Next
		nex.Next = mid.Next
		mid.Next = nex
		// printList(head)
	}
	// fmt.Println(mid)

	for mid = mid.Next; mid != nil; {
		if head.Val != mid.Val {
			return false
		}
		head = head.Next
		mid = mid.Next
	}

	// fmt.Println(mid)

	return true
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484822&idx=1&sn=44742c9a3557038c8da7150100d94db9&chksm=9bd7fb9eaca0728876e1146306a09f5453bcd5c35c4a264304ea6189faa83ec12a00322f0246&scene=21#wechat_redirect

/*
方法一: 递归

时间复杂度：
空间复杂度：
*/
func isPalindromeP1(head *ListNode) bool {
	left := head
	var reserve func(right *ListNode) bool
	reserve = func(right *ListNode) bool {
		if right == nil {
			return true
		}
		res := reserve(right.Next)
		res = res && right.Val == left.Val
		left = left.Next
		return res
	}

	return reserve(head)
}

// --- 官方

/*
方法一: 将值复制到数组中后用双指针法

时间复杂度：O(n)，其中 n 指的是链表的元素个数。
空间复杂度：O(n)，其中 n 指的是链表的元素个数，我们使用了一个数组列表存放链表的元素值。
*/
func isPalindromeO1(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

/*
方法二: 递归

时间复杂度：O(n)，其中 n 指的是链表的大小。
空间复杂度：O(n)，其中 n 指的是链表的大小。
*/
func isPalindromeO2(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

/*
方法三: 快慢指针

时间复杂度：O(n)，其中 n 指的是链表的大小。
空间复杂度：O(n)，其中 n 指的是链表的大小。
*/
func isPalindromeO3(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}
