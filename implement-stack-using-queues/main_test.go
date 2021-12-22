package implement_stack_using_queues

import (
	"testing"
)

type MyStackInterface interface {
	Push(x int)
	Pop() int
	Top() int
	Empty() bool
}

func TestMyStack(t *testing.T) {
	m := ConstructorS1()
	testEx1(&m, t)
}

func testEx1(s MyStackInterface, t *testing.T) {
	s.Push(1)
	s.Push(2)

	if got := s.Top(); got != 2 {
		t.Errorf("s.Top() = %v, want %v", got, 2)
	}

	if got := s.Pop(); got != 2 {
		t.Errorf("s.Pop() = %v, want %v", got, 2)
	}

	if got := s.Empty(); got != false {
		t.Errorf("s.Empty() = %v, want %v", got, false)
	}
}
