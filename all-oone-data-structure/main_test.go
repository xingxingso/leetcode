package all_oone_data_structure

import "testing"

type AllOneInterface interface {
	Inc(key string)
	Dec(key string)
	GetMaxKey() string
	GetMinKey() string
}

func TestAllOne(t *testing.T) {
	m := Constructor()
	testEx1(&m, t)
}

func testEx1(s AllOneInterface, t *testing.T) {
	s.Inc("hello")
	s.Inc("hello")

	if got, want := s.GetMaxKey(), "hello"; got != want {
		t.Errorf("s.GetMaxKey() = %v, want %v", got, want)
	}

	if got, want := s.GetMinKey(), "hello"; got != want {
		t.Errorf("s.GetMinKey() = %v, want %v", got, want)
	}

	s.Inc("leet")

	if got, want := s.GetMaxKey(), "hello"; got != want {
		t.Errorf("s.GetMaxKey() = %v, want %v", got, want)
	}

	if got, want := s.GetMinKey(), "leet"; got != want {
		t.Errorf("s.GetMinKey() = %v, want %v", got, want)
	}
}
