package reverse_linked_list_ii

import (
	"reflect"
	"testing"
)

type args struct {
	head  *ListNode
	left  int
	right int
}

func Test_reverseBetween(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetweenS2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetweenS2(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetweenS2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetweenS3(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetweenS3(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetweenS3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetweenO1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetweenO1(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_reverseBetweenO1() = %v, want %v", got, tt.want)
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

func getTests() []struct {
	name string
	args args
	want *ListNode
} {
	var tests = []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "equal0",
			args: args{
				head:  getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: getListNodeBySlice([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "equal1",
			args: args{
				head:  getListNodeBySlice([]int{5}),
				left:  1,
				right: 1,
			},
			want: getListNodeBySlice([]int{5}),
		},
		{
			name: "equal2",
			args: args{
				head:  getListNodeBySlice([]int{3, 5}),
				left:  1,
				right: 2,
			},
			want: getListNodeBySlice([]int{5, 3}),
		},
		{
			name: "equal3",
			args: args{
				head:  getListNodeBySlice([]int{1, 2, 3}),
				left:  3,
				right: 3,
			},
			want: getListNodeBySlice([]int{1, 2, 3}),
		},
	}
	return tests
}
