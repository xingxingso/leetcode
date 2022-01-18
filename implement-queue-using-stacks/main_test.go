package implement_queue_using_stacks

import "testing"

type MyQueueInterface interface {
	Push(x int)
	Pop() int
	Peek() int
	Empty() bool
}

func TestMyQueue(t *testing.T) {
	m := Constructor()
	testEx1(&m, t)
}

func testEx1(s MyQueueInterface, t *testing.T) {
	s.Push(1)
	s.Push(2)

	if got := s.Peek(); got != 1 {
		t.Errorf("s.Top() = %v, want %v", got, 1)
	}

	if got := s.Pop(); got != 1 {
		t.Errorf("s.Pop() = %v, want %v", got, 1)
	}

	if got := s.Empty(); got != false {
		t.Errorf("s.Empty() = %v, want %v", got, false)
	}
}
