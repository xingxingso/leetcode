package linked_list_cycle

import "testing"

type args struct {
	head *ListNode
}

func getTests() []struct {
	name string
	args args
	want bool
} {
	return []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				head: ex1(),
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				head: ex2(),
			},
			want: true,
		},
		{
			name: "ex3",
			args: args{
				head: ex3(),
			},
			want: false,
		},
	}
}

func Test_hasCycle(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
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
