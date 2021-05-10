package public

type Stack interface {
	IsEmpty() bool
	Pop() int
	Push(val int)
	Peek() int
}

type StackList struct {
	head *ListNode
}

func NewStackList() *StackList {
	return &StackList{
		head: &ListNode{},
	}
}

func (stack *StackList) IsEmpty() bool {
	return stack.head.Next == nil
}

func (stack *StackList) Pop() int {
	if stack.head.Next == nil {
		return -1
	}

	val := stack.head.Next.Val
	stack.head.Next = stack.head.Next.Next
	return val
}

func (stack *StackList) Push(val int) {
	node := &ListNode{
		Val: val,
	}

	node.Next = stack.head.Next
	stack.head.Next = node
}

func (stack *StackList) Peek() int {
	if stack.head.Next == nil {
		return -1
	}

	return stack.head.Next.Val
}

type StackSlice []int

func NewStackSlice() *StackSlice {
	return &StackSlice{}
}

func (s *StackSlice) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StackSlice) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return val
}

func (s *StackSlice) Push(val int) {
	*s = append(*s, val)
}

func (s *StackSlice) Peek() int {
	if s.IsEmpty() {
		return 0
	}

	return (*s)[len(*s)-1]
}
