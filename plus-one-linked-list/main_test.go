package plus_one_linked_list

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
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
				head: getListNodeBySlice([]int{1, 2, 3}),
			},
			want: getListNodeBySlice([]int{1, 2, 4}),
		},
		{
			name: "ex2",
			args: args{
				head: getListNodeBySlice([]int{9, 9, 9}),
			},
			want: getListNodeBySlice([]int{1, 0, 0, 0}),
		},
		{
			name: "ex3",
			args: args{
				head: getListNodeBySlice([]int{0}),
			},
			want: getListNodeBySlice([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
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
