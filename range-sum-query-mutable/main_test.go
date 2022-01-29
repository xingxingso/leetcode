package range_sum_query_mutable

import "testing"

type NumArrayInterface interface {
	Update(index int, val int)
	SumRange(i, j int) int
}

func TestNumArray(t *testing.T) {
	s1 := ConstructorS1([]int{1, 3, 5})
	testEx1(&s1, t)
	s1 = ConstructorS1([]int{7, 2, 7, 2, 0})
	testEx2(&s1, t)

	// 方法二
	s2 := Constructor([]int{1, 3, 5})
	testEx1(&s2, t)
	s2 = Constructor([]int{7, 2, 7, 2, 0})
	testEx2(&s2, t)
}

func testEx1(s NumArrayInterface, t *testing.T) {
	if got := s.SumRange(0, 2); got != 9 {
		t.Errorf("s.SumRange() = %v, want %v", got, 9)
	}

	s.Update(1, 2)

	if got := s.SumRange(0, 2); got != 8 {
		t.Errorf("s.SumRange() = %v, want %v", got, 8)
	}
}

func testEx2(s NumArrayInterface, t *testing.T) {
	s.Update(4, 6)
	s.Update(0, 2)
	s.Update(0, 9)

	if got := s.SumRange(4, 4); got != 6 {
		t.Errorf("s.SumRange() = %v, want %v", got, 6)
	}

	s.Update(3, 8)

	if got := s.SumRange(0, 4); got != 32 {
		t.Errorf("s.SumRange() = %v, want %v", got, 32)
	}

	s.Update(4, 1)

	if got := s.SumRange(0, 3); got != 26 {
		t.Errorf("s.SumRange() = %v, want %v", got, 26)
	}
	if got := s.SumRange(0, 4); got != 27 {
		t.Errorf("s.SumRange() = %v, want %v", got, 27)
	}

	s.Update(0, 4)
}
