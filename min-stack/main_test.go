package min_stack

import "testing"

type MinStackInterface interface {
	Push(x int)
	Pop()
	Top() int
	GetMin() int
}

func TestMinStack(t *testing.T) {
	m := Constructor()
	testEx1(&m, t)

	p1 := ConstructorP1()
	testEx1(&p1, t)
}

func testEx1(s MinStackInterface, t *testing.T) {
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.Push(-3)

	if got := s.GetMin(); got != -3 {
		t.Errorf("s.GetMin() = %v, want %v", got, -3)
	}

	s.Pop()
	if got := s.GetMin(); got != -3 {
		t.Errorf("s.GetMin() = %v, want %v", got, -3)
	}
	s.Pop()

	if got := s.Top(); got != 0 {
		t.Errorf("s.Top() = %v, want %v", got, 0)
	}

	if got := s.GetMin(); got != -2 {
		t.Errorf("s.GetMin() = %v, want %v", got, -2)
	}
}
