package odd_even_linked_list

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
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
				head: GetListNodeBySlice([]int{1, 2, 3, 4, 5}),
			},
			want: GetListNodeBySlice([]int{1, 3, 5, 2, 4}),
		},
		{
			name: "ex2",
			args: args{
				head: GetListNodeBySlice([]int{2, 1, 3, 5, 6, 4, 7}),
			},
			want: GetListNodeBySlice([]int{2, 3, 6, 7, 1, 5, 4}),
		},
		{
			name: "ex3",
			args: args{
				head: GetListNodeBySlice([]int{}),
			},
			want: GetListNodeBySlice([]int{}),
		},
		{
			name: "ex4",
			args: args{
				head: GetListNodeBySlice([]int{1}),
			},
			want: GetListNodeBySlice([]int{1}),
		},
		{
			name: "ex5",
			args: args{
				head: GetListNodeBySlice([]int{1, 2, 3, 4}),
			},
			want: GetListNodeBySlice([]int{1, 3, 2, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
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
