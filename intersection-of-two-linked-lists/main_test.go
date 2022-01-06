package intersection_of_two_linked_lists

import (
	"reflect"
	"testing"
)

type args struct {
	headA     *ListNode
	headB     *ListNode
	intersect *ListNode
}

func Test_getIntersectionNode(t *testing.T) {
	argArr := createArgArr()

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex1",
			args: args{
				headA: argArr[0].headA,
				headB: argArr[0].headB,
			},
			want: argArr[0].intersect,
		},
		{
			name: "ex2",
			args: args{
				headA: argArr[1].headA,
				headB: argArr[1].headB,
			},
			want: argArr[1].intersect,
		},
		{
			name: "ex3",
			args: args{
				headA: argArr[2].headA,
				headB: argArr[2].headB,
			},
			want: argArr[2].intersect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}

			if got := getIntersectionNodeO1(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNodeO1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createArgArr() []*args {
	res := make([]*args, 0)
	var listA, listB *ListNode
	var arg *args

	// 测试用例1
	listA = GetListNodeBySlice([]int{4, 1, 8, 4, 5})
	listB = GetListNodeBySlice([]int{5, 6, 1}) // 8 4 5
	listB.Next.Next.Next = listA.Next.Next
	arg = &args{
		headA:     listA,
		headB:     listB,
		intersect: listA.Next.Next,
	}
	res = append(res, arg)

	// 测试用例2
	listA = GetListNodeBySlice([]int{1, 9, 1, 2, 4})
	listB = GetListNodeBySlice([]int{3}) // 2 4
	listB.Next = listA.Next.Next.Next
	arg = &args{
		headA:     listA,
		headB:     listB,
		intersect: listB.Next,
	}
	res = append(res, arg)

	// 测试用例3
	listA = GetListNodeBySlice([]int{2, 6, 4})
	listB = GetListNodeBySlice([]int{1, 5})
	arg = &args{
		headA:     listA,
		headB:     listB,
		intersect: nil,
	}
	res = append(res, arg)

	return res
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
