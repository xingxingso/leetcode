package linked_list_cycle_ii

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	node1 := ex1()
	node2 := ex2()
	node3 := ex3()
	tests := []struct {
		name string
		args args
		want *ListNode
	}{

		{
			name: "ex1",
			args: args{
				head: node1,
			},
			want: node1.Next,
		},
		{
			name: "ex2",
			args: args{
				head: node2,
			},
			want: node2,
		},
		{
			name: "ex3",
			args: args{
				head: node3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ex1() *ListNode {
	head := &ListNode{
		Val: 3,
	}
	temp := head

	posNode := &ListNode{Val: 2}
	temp.Next = posNode
	temp = temp.Next

	temp.Next = &ListNode{Val: 0}
	temp = temp.Next

	temp.Next = &ListNode{Val: -4}
	temp = temp.Next

	temp.Next = posNode

	return head
}

func ex2() *ListNode {
	head := &ListNode{
		Val: 1,
	}
	temp := head
	posNode := head

	temp.Next = &ListNode{Val: 2}
	temp = temp.Next

	temp.Next = posNode

	return head
}

func ex3() *ListNode {
	head := &ListNode{
		Val: 1,
	}
	return head
}
