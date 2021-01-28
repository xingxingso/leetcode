package public

type StackSlice struct {
	stack []int
}

func NewStackSlice() *StackSlice {
	return &StackSlice{}
}

func (s *StackSlice) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *StackSlice) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]

	return val
}

func (s *StackSlice) Push(val int) {
	s.stack = append(s.stack, val)
}

func (s *StackSlice) Peek() int {
	if s.IsEmpty() {
		return 0
	}

	return s.stack[len(s.stack)-1]
}
