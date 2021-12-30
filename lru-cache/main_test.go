package lru_cache

import (
	"testing"
)

type LRUInterface interface {
	Get(key int) int
	Put(key int, value int)
}

func TestLRUCache(t *testing.T) {
	s1 := ConstructorS1(2)
	testEx1(&s1, t)

	s2 := Constructor(2)
	testEx1(&s2, t)

	l := ConstructorO1(2)
	testEx1(&l, t)
}

func testEx1(l LRUInterface, t *testing.T) {
	l.Put(1, 1)
	l.Put(2, 2)

	if got := l.Get(1); got != 1 {
		t.Errorf("Get() = %v, want %v", got, 1)
	}

	l.Put(3, 3)

	if got := l.Get(2); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	l.Put(4, 4)

	if got := l.Get(1); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	if got := l.Get(3); got != 3 {
		t.Errorf("Get() = %v, want %v", got, 3)
	}

	if got := l.Get(4); got != 4 {
		t.Errorf("Get() = %v, want %v", got, 4)
	}
}
