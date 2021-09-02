/*
Package merge_k_sorted_lists
https://leetcode-cn.com/problems/merge-k-sorted-lists/

23. 合并K个升序链表

给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

提示：
	k == lists.length
	0 <= k <= 10^4
	0 <= lists[i].length <= 500
	-10^4 <= lists[i][j] <= 10^4
	lists[i] 按 升序 排列
	lists[i].length 的总和不超过 10^4

*/
package merge_k_sorted_lists

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
func mergeKLists(lists []*ListNode) *ListNode {
	ans := &ListNode{}
	head := ans
	for {
		minIdx := -1
		for k, v := range lists {
			if v != nil {
				minIdx = k
				break
			}
		}
		if minIdx == -1 {
			break
		}
		for i := minIdx + 1; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if lists[i].Val < lists[minIdx].Val {
				minIdx = i
			}
		}
		head.Next = lists[minIdx]
		head = head.Next
		lists[minIdx] = lists[minIdx].Next
	}
	return ans.Next
}
