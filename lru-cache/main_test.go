package lru_cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	l := Constructor(2)

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

func TestLRUCacheS1(t *testing.T) {
	l := ConstructorS1(2)

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
