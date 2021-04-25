package reverse_linked_list

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "equal0",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
			},
			want: getListNodeBySlice([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !isListNodeValEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}

	tests2 := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "equal",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
			},
			want: getListNodeBySlice([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListO1(tt.args.head); !isListNodeValEqual(got, tt.want) {
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
