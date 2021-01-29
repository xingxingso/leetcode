package public

import "reflect"

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
loop:
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				continue loop
			}
		}
		return false
	}

	return true
}
