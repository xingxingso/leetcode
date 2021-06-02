package remove_duplicates_from_sorted_list

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
	return []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				head: getListNodeBySlice([]int{1, 1, 2}),
			},
			want: getListNodeBySlice([]int{1, 2}),
		},
		{
			name: "ex2",
			args: args{
				head: getListNodeBySlice([]int{1, 1, 2, 3, 3}),
			},
			want: getListNodeBySlice([]int{1, 2, 3}),
		},
		{
			name: "ex3",
			args: args{
				head: getListNodeBySlice([]int{}),
			},
			want: getListNodeBySlice([]int{}),
		},
	}
}

func Test_deleteDuplicates(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicatesS1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesS1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicatesS1() = %v, want %v", got, tt.want)
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
