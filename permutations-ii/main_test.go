package permutations_ii

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUniqueO1(tt.args.nums); !intSliceSliceEqual2(got, tt.want) {
				t.Errorf("permuteUniqueO1() = %v, want %v", got, tt.want)
			}
		})
	}
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
