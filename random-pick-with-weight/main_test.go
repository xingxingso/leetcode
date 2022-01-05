package random_pick_with_weight

import (
	"fmt"
	"testing"
)

type SolutionInterface interface {
	PickIndex() int
}

func TestSolution(t *testing.T) {
	type args struct {
		w []int
	}
	tests := []struct {
		name string
		args args
		want map[int]float64
	}{
		{
			name: "ex1",
			args: args{
				w: []int{1},
			},
			want: map[int]float64{0: 1.0},
		},
		{
			name: "ex2",
			args: args{
				w: []int{1, 3},
			},
			want: map[int]float64{0: 0.25, 1: 0.75},
		},
		{
			name: "ex3",
			args: args{
				w: []int{100},
			},
			want: map[int]float64{0: 1},
		},
	}

	for _, test := range tests {
		s := Constructor(test.args.w)
		ex1(&s, test.want, t)
	}
}

func ex1(s SolutionInterface, want map[int]float64, t *testing.T) {
	total := 100000
	count := make(map[int]int)
	for i := 0; i < total; i++ {
		count[s.PickIndex()]++
	}
	got := make(map[int]float64)

	for k, v := range count {
		got[k] = float64(v) / float64(total)
	}
	fmt.Println(got)

	for k, v := range got {
		if _, ok := want[k]; !ok {
			t.Errorf("num %v not want, got ratio %v", k, v)
			continue
		}

		if abs(want[k]-v) > 0.01 {
			t.Errorf("num %v want %v, got %v", k, want[k], v)
		}
	}
}

func abs(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}
