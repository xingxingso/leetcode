package reverse_nodes_in_k_group

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "equal",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: getListNodeBySlice([]int{2, 1, 4, 3, 5}),
		},
		{
			name: "equal",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: getListNodeBySlice([]int{3, 2, 1, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
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
