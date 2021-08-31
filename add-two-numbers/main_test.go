package add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				l1: getListNodeBySlice([]int{2, 4, 3}),
				l2: getListNodeBySlice([]int{5, 6, 4}),
			},
			want: getListNodeBySlice([]int{7, 0, 8}),
		},
		{
			name: "ex2",
			args: args{
				l1: getListNodeBySlice([]int{0}),
				l2: getListNodeBySlice([]int{0}),
			},
			want: getListNodeBySlice([]int{0}),
		},
		{
			name: "ex3",
			args: args{
				l1: getListNodeBySlice([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: getListNodeBySlice([]int{9, 9, 9, 9}),
			},
			want: getListNodeBySlice([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
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
