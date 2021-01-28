package public

type Stack struct {
	head *ListNode
}

func NewStack() *Stack {
	return &Stack{
		head: &ListNode{},
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.head.Next == nil
}

func (stack *Stack) Pop() int {
	if stack.head.Next == nil {
		return -1
	}

	val := stack.head.Next.Val
	stack.head.Next = stack.head.Next.Next
	return val
}

func (stack *Stack) Push(val int) {
	node := &ListNode{
		Val: val,
	}

	node.Next = stack.head.Next
	stack.head.Next = node
}

func (stack *Stack) Peek() int {
	if stack.head.Next == nil {
		return -1
	}

	return stack.head.Next.Val
}
