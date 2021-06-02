package public

import "fmt"

/*
ListNode
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleList struct {
	Key, Val  int
	Pre, Next *DoubleList
}

func getListNodeBySlice(s []int) *ListNode {
	head := &ListNode{Val: 0}
	tmp := head

	for _, v := range s {
		node := &ListNode{Val: v}
		head.Next = node
		head = node
	}
	return tmp.Next
}

func isListNodeValEqual(node1, node2 *ListNode) bool {
	for node1 != nil {
		if node2 == nil || node1.Val != node2.Val {
			return false
		}

		node1 = node1.Next
		node2 = node2.Next
	}

	return node2 == nil
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}
