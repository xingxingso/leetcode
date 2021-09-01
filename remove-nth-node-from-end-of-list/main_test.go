package remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: getListNodeBySlice([]int{1, 2, 3, 5}),
		},
		{
			name: "ex2",
			args: args{
				head: getListNodeBySlice([]int{1}),
				n:    1,
			},
			want: getListNodeBySlice([]int{}),
		},
		{
			name: "ex3",
			args: args{
				head: getListNodeBySlice([]int{1, 2}),
				n:    1,
			},
			want: getListNodeBySlice([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
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
