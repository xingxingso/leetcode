package linked_list_random_node

import (
	"fmt"
	"testing"
)

// 未验证随机性
func TestSolutionS1(t *testing.T) {
	// 初始化一个单链表 [1,2,3].
	head := getListNodeBySlice([]int{1, 2, 3})
	s := ConstructorS1(head)
	// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
	count := make([]int, 3)
	for i := 0; i < 10000; i++ {
		n := s.GetRandom()
		if n != 1 && n != 2 && n != 3 {
			t.Errorf("GetRandom=%v", n)
			break
		}
		count[n-1]++
	}
	for i, v := range count {
		if v == 0 {
			t.Errorf("this number: %v not got", i+1)
		}
	}
	fmt.Println(count)
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
