package palindrome_linked_list

import "testing"

type args struct {
	head *ListNode
}

func getTest() []struct {
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
			name: "equal0",
			args: args{
				head: getListNodeBySlice([]int{1, 2}),
			},
			want: false,
		},
		{
			name: "equal1",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 2, 1}),
			},
			want: true,
		},
		{
			name: "equal2",
			args: args{
				head: getListNodeBySlice([]int{1, 2, 1}),
			},
			want: true,
		},
		{
			name: "equal3",
			args: args{
				head: getListNodeBySlice([]int{1, 3, 2, 4, 3, 2, 1}),
			},
			want: false,
		},
	}
}

func Test_isPalindrome(t *testing.T) {
	for _, tt := range getTest() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeP1(t *testing.T) {
	for _, tt := range getTest() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeP1(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeP1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeO1(t *testing.T) {
	for _, tt := range getTest() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeO1(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeO1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeO2(t *testing.T) {
	for _, tt := range getTest() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeO2(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeO2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeO3(t *testing.T) {
	for _, tt := range getTest() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeO3(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeO3() = %v, want %v", got, tt.want)
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
