package reverse_nodes_in_k_group

import (
	"reflect"
	"testing"
)

type args struct {
	head *ListNode
	k    int
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
			name: "ex0",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: getListNodeBySlice([]int{2, 1, 4, 3, 5}),
		},
		{
			name: "ex1",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: getListNodeBySlice([]int{3, 2, 1, 4, 5}),
		},
		{
			name: "ex2",
			args: args{
				head: getListNodeBySlice([]int{1, 2}),
				k:    2,
			},
			want: getListNodeBySlice([]int{2, 1}),
		},
	}
}

func Test_reverseKGroup(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroupS1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroupS1(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroupS1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroupP1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroupP1(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroupP1() = %v, want %v", got, tt.want)
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
