package shuffle_an_array

import (
	"fmt"
	"reflect"
	"testing"
)

type solutionInterface interface {
	Reset() []int
	Shuffle() []int
}

func TestSolution(t *testing.T) {
	m := Constructor([]int{1, 2, 3})
	testEx1(&m, t)

	m2 := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	testEx2(&m2, t)
}

func testEx1(s solutionInterface, t *testing.T) {
	var got []int
	got = s.Shuffle()
	fmt.Println(got)
	if got, want := s.Reset(), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("s.Reset() = %v, want %v", got, want)
	}
	got = s.Shuffle()
	fmt.Println(got)
}

func testEx2(s solutionInterface, t *testing.T) {
	var got []int
	got = s.Shuffle()
	fmt.Println(got)
	if got, want := s.Reset(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}; !reflect.DeepEqual(got, want) {
		t.Errorf("s.Reset() = %v, want %v", got, want)
	}
	got = s.Shuffle()
	fmt.Println(got)
}
