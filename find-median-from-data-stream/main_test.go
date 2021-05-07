package find_median_from_data_stream

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	m := Constructor()
	testEx1(&m, t)
	m2 := Constructor()
	testEx2(&m2, t)
}

type MF interface {
	AddNum(num int)
	FindMedian() float64
}

func testEx1(mf MF, t *testing.T) {
	mf.AddNum(1)
	mf.AddNum(2)

	if got := mf.FindMedian(); got != 1.5 {
		t.Errorf("FindMedian() = %v, want %v", got, 1.5)
	}

	mf.AddNum(3)
	if got := mf.FindMedian(); got != 2 {
		t.Errorf("FindMedian() = %v, want %v", got, 2)
	}
}

func testEx2(mf MF, t *testing.T) {
	mf.AddNum(-1)
	if got := mf.FindMedian(); got != -1.0 {
		t.Errorf("FindMedian() = %v, want %v", got, -1.0)
	}

	mf.AddNum(-2)
	if got := mf.FindMedian(); got != -1.5 {
		t.Errorf("FindMedian() = %v, want %v", got, -1.5)
	}

	mf.AddNum(-3)
	if got := mf.FindMedian(); got != -2.0 {
		t.Errorf("FindMedian() = %v, want %v", got, -2.0)
	}

	mf.AddNum(-4)
	if got := mf.FindMedian(); got != -2.5 {
		t.Errorf("FindMedian() = %v, want %v", got, -2.5)
	}

	mf.AddNum(-5)
	if got := mf.FindMedian(); got != -3.0 {
		t.Errorf("FindMedian() = %v, want %v", got, -3.0)
	}
}
