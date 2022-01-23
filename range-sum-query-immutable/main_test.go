package range_sum_query_immutable

import "testing"

type NumArrayInterface interface {
	SumRange(i, j int) int
}

func TestNumArray(t *testing.T) {
	m := Constructor([]int{-2, 0, 3, -5, 2, -1})
	testEx1(&m, t)
}

func testEx1(s NumArrayInterface, t *testing.T) {
	if got := s.SumRange(0, 2); got != 1 {
		t.Errorf("s.SumRange() = %v, want %v", got, 1)
	}

	if got := s.SumRange(2, 5); got != -1 {
		t.Errorf("s.SumRange() = %v, want %v", got, -1)
	}

	if got := s.SumRange(0, 5); got != -3 {
		t.Errorf("s.SumRange() = %v, want %v", got, -3)
	}
}
