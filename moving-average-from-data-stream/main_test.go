package moving_average_from_data_stream

import (
	"testing"
)

func TestMovingAverage(t *testing.T) {
	m := Constructor(3)

	want := 1.0
	if got := m.Next(1); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	want = 5.5
	if got := m.Next(10); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	want = 4.666666666666667
	if got := m.Next(3); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	want = 6.0
	if got := m.Next(5); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}

	// test size 1
	m = Constructor(1)

	want = 1.0
	if got := m.Next(1); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	want = 10.0
	if got := m.Next(10); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	want = 3.0
	if got := m.Next(3); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	want = 5.0
	if got := m.Next(5); got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
}
