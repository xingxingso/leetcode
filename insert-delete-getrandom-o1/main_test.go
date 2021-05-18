package insert_delete_getrandom_o1

import (
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	// s := Constructor()
	// testEx1(&s, t)

	s2 := Constructor()
	testEx2(&s2, t)
}

type RandomizedSetInterface interface {
	Insert(val int) bool
	Remove(val int) bool
	GetRandom() int
}

func testEx1(s RandomizedSetInterface, t *testing.T) {

	// 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	if got := s.Insert(1); got != true {
		t.Errorf("Insert() = %v, want %v", got, true)
	}

	// 返回 false ，表示集合中不存在 2 。
	if got := s.Remove(2); got != false {
		t.Errorf("Remove() = %v, want %v", got, false)
	}

	// 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	if got := s.Insert(2); got != true {
		t.Errorf("Insert() = %v, want %v", got, true)
	}

	// getRandom 应随机返回 1 或 2 。
	if got := s.GetRandom(); got != 1 && got != 2 {
		t.Errorf("GetRandom() = %v, want %v or %v", got, 1, 2)
	}

	// 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	if got := s.Remove(1); got != true {
		t.Errorf("Remove() = %v, want %v", got, true)
	}

	// 2 已在集合中，所以返回 false 。
	if got := s.Insert(2); got != false {
		t.Errorf("Insert() = %v, want %v", got, false)
	}

	// 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
	if got := s.GetRandom(); got != 2 {
		t.Errorf("GetRandom() = %v, want %v", got, 2)
	}
}

func testEx2(s RandomizedSetInterface, t *testing.T) {
	if got := s.Insert(0); got != true {
		t.Errorf("Insert() = %v, want %v", got, true)
	}
	if got := s.Insert(1); got != true {
		t.Errorf("Insert() = %v, want %v", got, true)
	}
	if got := s.Remove(0); got != true {
		t.Errorf("Remove() = %v, want %v", got, true)
	}
	if got := s.Insert(2); got != true {
		t.Errorf("Insert() = %v, want %v", got, true)
	}
	if got := s.Remove(1); got != true {
		t.Errorf("Remove() = %v, want %v", got, true)
	}
	if got := s.GetRandom(); got != 2 {
		t.Errorf("GetRandom() = %v, want %v", got, 2)
	}
}
