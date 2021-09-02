package merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				lists: []*ListNode{
					getListNodeBySlice([]int{1, 4, 5}),
					getListNodeBySlice([]int{1, 3, 4}),
					getListNodeBySlice([]int{2, 6}),
				},
			},
			want: getListNodeBySlice([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "ex2",
			args: args{
				lists: []*ListNode{},
			},
			want: getListNodeBySlice([]int{}),
		},
		{
			name: "ex3",
			args: args{
				lists: []*ListNode{getListNodeBySlice([]int{})},
			},
			want: getListNodeBySlice([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
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
