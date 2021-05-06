package find_median_from_data_stream

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	m := Constructor()
	testEx1(&m, t)
}

type MF interface {
	AddNum(num int)
	FindMedian() float64
}

func testEx1(mf MF, t *testing.T) {
	mf.AddNum(1)
	mf.AddNum(2)
	mf.AddNum(5)
	mf.AddNum(3)
	mf.AddNum(7)
	mf.AddNum(9)
	mf.AddNum(4)
	mf.AddNum(8)
	mf.AddNum(6)

	if got := mf.FindMedian(); got != 1.5 {
		t.Errorf("Get() = %v, want %v", got, 1.5)
	}
	mf.AddNum(3)
	if got := mf.FindMedian(); got != 2 {
		t.Errorf("Get() = %v, want %v", got, 2)
	}
}
