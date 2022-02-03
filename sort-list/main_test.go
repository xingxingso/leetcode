package sort_list

import (
	"reflect"
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
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				head: GetListNodeBySlice([]int{4, 2, 1, 3}),
			},
			want: GetListNodeBySlice([]int{1, 2, 3, 4}),
		},
		{
			name: "ex2",
			args: args{
				head: GetListNodeBySlice([]int{-1, 5, 3, 4, 0}),
			},
			want: GetListNodeBySlice([]int{-1, 0, 3, 4, 5}),
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
				head: GetListNodeBySlice([]int{2, 1}),
			},
			want: GetListNodeBySlice([]int{1, 2}),
		},
	}
	return tests
}

func Test_sortList(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortListO1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortListO1() = %v, want %v", got, tt.want)
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
