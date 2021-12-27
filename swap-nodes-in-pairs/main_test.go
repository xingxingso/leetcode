package swap_nodes_in_pairs

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				head: GetListNodeBySlice([]int{1, 2, 3, 4}),
			},
			want: GetListNodeBySlice([]int{2, 1, 4, 3}),
		},
		{
			name: "ex2",
			args: args{
				head: GetListNodeBySlice([]int{}),
			},
			want: GetListNodeBySlice([]int{}),
		},
		{
			name: "ex3",
			args: args{
				head: GetListNodeBySlice([]int{1}),
			},
			want: GetListNodeBySlice([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

// GetListNodeBySlice 通过 slice 构建 list
func GetListNodeBySlice(s []int) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for _, v := range s {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	return dummy.Next
}

// IsListNodeValEqual 判断两个链表是否一致
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
