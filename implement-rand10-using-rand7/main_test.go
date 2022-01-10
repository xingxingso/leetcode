package implement_rand10_using_rand7

import (
	"fmt"
	"testing"
)

func Test_rand10(t *testing.T) {
	n := 100000
	got := make([]int, 10)
	for i := 0; i < n; i++ {
		x := rand10()
		if x > 10 || x < 1 {
			t.Errorf("rand10() got: %v", x)
			break
		}
		got[x-1]++
	}

	fmt.Println(got)

	for i := 0; i < 10; i++ {
		diff := abs(got[i] - n/10)
		if float64(diff)/float64(n/10) > 0.03 {
			t.Errorf("rand10() %d diff:%v", i+1, diff)
		}
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
