package merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				list1: GetListNodeBySlice([]int{1, 2, 4}),
				list2: GetListNodeBySlice([]int{1, 3, 4}),
			},
			want: GetListNodeBySlice([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "ex2",
			args: args{
				list1: GetListNodeBySlice([]int{}),
				list2: GetListNodeBySlice([]int{}),
			},
			want: GetListNodeBySlice([]int{}),
		},
		{
			name: "ex3",
			args: args{
				list1: GetListNodeBySlice([]int{}),
				list2: GetListNodeBySlice([]int{0}),
			},
			want: GetListNodeBySlice([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
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
