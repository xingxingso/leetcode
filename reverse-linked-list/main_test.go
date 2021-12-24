package reverse_linked_list

import (
	"testing"
)

type args struct {
	head *ListNode
}

func getTests() []struct {
	name string
	args args
	want *ListNode
} {
	return []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
			},
			want: getListNodeBySlice([]int{5, 4, 3, 2, 1}),
		},
	}
}

func Test_reverseList(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !IsListNodeValEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListO1(tt.args.head); !IsListNodeValEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
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

func IsListNodeValEqual(node1, node2 *ListNode) bool {
	for node1 != nil {
		if node2 == nil || node1.Val != node2.Val {
			return false
		}

		node1 = node1.Next
		node2 = node2.Next
	}

	return node2 == nil
}
