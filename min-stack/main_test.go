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

	if got, want := s.GetMin(), -3; got != want {
		t.Errorf("s.GetMin() = %v, want %v", got, want)
	}

	s.Pop()
	if got, want := s.GetMin(), -3; got != want {
		t.Errorf("s.GetMin() = %v, want %v", got, want)
	}
	s.Pop()

	if got, want := s.Top(), 0; got != want {
		t.Errorf("s.Top() = %v, want %v", got, want)
	}

	if got, want := s.GetMin(), -2; got != want {
		t.Errorf("s.GetMin() = %v, want %v", got, want)
	}
}
