package max_stack

import (
	"testing"
)

func TestMaxStack(t *testing.T) {
	stk := Constructor()

	stk.Push(5)
	stk.Push(1)
	stk.Push(5)

	if got := stk.Top(); got != 5 {
		t.Errorf("Top() = %v, want %v", got, 5)
	}

	if got := stk.PopMax(); got != 5 {
		t.Errorf("PopMax() = %v, want %v", got, 5)
	}

	if got := stk.Top(); got != 1 {
		t.Errorf("Top() = %v, want %v", got, 1)
	}

	if got := stk.PeekMax(); got != 5 {
		t.Errorf("PeekMax() = %v, want %v", got, 5)
	}

	if got := stk.Pop(); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}

	if got := stk.Top(); got != 5 {
		t.Errorf("Top() = %v, want %v", got, 5)
	}
}

func TestMaxStackO1(t *testing.T) {
	stk := ConstructorO1()

	stk.Push(5)
	stk.Push(1)
	stk.Push(5)

	if got := stk.Top(); got != 5 {
		t.Errorf("Top() = %v, want %v", got, 5)
	}

	if got := stk.PopMax(); got != 5 {
		t.Errorf("PopMax() = %v, want %v", got, 5)
	}

	if got := stk.Top(); got != 1 {
		t.Errorf("Top() = %v, want %v", got, 1)
	}

	if got := stk.PeekMax(); got != 5 {
		t.Errorf("PeekMax() = %v, want %v", got, 5)
	}

	if got := stk.Pop(); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}

	if got := stk.Top(); got != 5 {
		t.Errorf("Top() = %v, want %v", got, 5)
	}
}

func TestMaxStackO2(t *testing.T) {
	stk := ConstructorO2()

	stk.Push(5)
	stk.Push(1)
	stk.Push(5)

	if got := stk.Top(); got != 5 {
		t.Errorf("Top() = %v, want %v", got, 5)
	}

	if got := stk.PopMax(); got != 5 {
		t.Errorf("PopMax() = %v, want %v", got, 5)
	}

	if got := stk.Top(); got != 1 {
		t.Errorf("Top() = %v, want %v", got, 1)
	}

	if got := stk.PeekMax(); got != 5 {
		t.Errorf("PeekMax() = %v, want %v", got, 5)
	}

	if got := stk.Pop(); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}

	if got := stk.Top(); got != 5 {
		t.Errorf("Top() = %v, want %v", got, 5)
	}
}
