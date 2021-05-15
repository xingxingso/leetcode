package middle_of_the_linked_list

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
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
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
			},
			want: getListNodeBySlice([]int{3, 4, 5}),
		},
		{
			name: "ex2",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5, 6}),
			},
			want: getListNodeBySlice([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
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
