package range_sum_query_2d_immutable

import "testing"

type NumMatrixInterface interface {
	SumRegion(row1 int, col1 int, row2 int, col2 int) int
}

func TestNumMatrix(t *testing.T) {
	m := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	testEx1(&m, t)

	s2 := ConstructorS2([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	testEx1(&s2, t)
}

func testEx1(s NumMatrixInterface, t *testing.T) {
	if got := s.SumRegion(2, 1, 4, 3); got != 8 {
		t.Errorf("s.SumRegion() = %v, want %v", got, 8)
	}

	if got := s.SumRegion(1, 1, 2, 2); got != 11 {
		t.Errorf("s.SumRegion() = %v, want %v", got, 11)
	}

	if got := s.SumRegion(1, 2, 2, 4); got != 12 {
		t.Errorf("s.SumRegion() = %v, want %v", got, 12)
	}
}
