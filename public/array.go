package public

import (
	"reflect"
	"sort"
)

func sum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}

	return ans
}

func inArray(nums []int, n int) bool {
	for _, v := range nums {
		if v == n {
			return true
		}
	}
	return false
}

func copy2(dst, src [][]int) {
	for key, value := range src {
		dst[key] = make([]int, len(value))
		copy(dst[key], value)
	}
}

func stringSliceSliceEqual(v1, v2 [][]string) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if reflect.DeepEqual(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	return true
}

// 是否含有相同元素，顺序无关, 其子元素要求顺序相同
// [[1,2]] != [[2,1]], [[2],[1]] == [[1],[2]]
func intSliceSliceEqual2(v1, v2 [][]int) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if len(vv1) != len(vv2) {
				continue
			}
			if reflect.DeepEqual(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	// loop 中 只能保证所有 v1 中元素 可在 v2 找到，当 v1 有重复元素时，v2 可能只有一个对应元素
loop2:
	for _, vv2 := range v2 {
		for _, vv1 := range v1 {
			if len(vv1) != len(vv2) {
				continue
			}
			if reflect.DeepEqual(vv1, vv2) {
				continue loop2
			}
		}

		return false
	}

	return true
}

// 二维数组是否相同， 不考虑元素顺序及子数组顺序 [[1,2]] == [[2,1]]
func intSliceSliceEqual(v1, v2 [][]int) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if len(vv1) != len(vv2) {
				continue
			}
			if sliceSameValue(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	return true
}

func sliceSameValue(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	t1, t2 := make([]int, len(s1)), make([]int, len(s2))
	copy(t1, s1)
	copy(t2, s2)
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	sort.Slice(t2, func(i, j int) bool {
		return t2[i] < t2[j]
	})
	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}
