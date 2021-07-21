package random_pick_index

import (
	"fmt"
	"testing"
)

// 未验证随机性
func TestSolutionS1(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3}
	s := ConstructorS1(nums)
	// pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
	count := make([]int, 3)
	for i := 0; i < 10000; i++ {
		n := s.Pick(3)
		if n >= len(nums) || n < 0 || nums[n] != 3 {
			t.Errorf("Pick=%v", n)
			break
		}
		count[n-2]++
	}
	fmt.Println(count)

	for i := 0; i < 10000; i++ {
		n := s.Pick(1)
		if n >= len(nums) || n < 0 || nums[n] != 0 {
			t.Errorf("Pick=%v", n)
			break
		}
	}
}
