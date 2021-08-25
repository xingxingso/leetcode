package pancake_sorting

import (
	"testing"
)

func Test_pancakeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				arr: []int{3, 2, 4, 1},
			},
			want: []int{4, 2, 4, 3},
		},
		{
			name: "ex2",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test pancakeSort
			temp := make([]int, len(tt.args.arr))
			copy(temp, tt.args.arr)
			if got := pancakeSort(tt.args.arr); !testPancakeSort(temp, got) {
				t.Errorf("pancakeSort() = %v, want %v", got, tt.want)
			}
		})
	}

}

func testPancakeSort(arr, sortArr []int) bool {
	if len(sortArr) > 10*len(arr) {
		return false
	}

	for _, v := range sortArr {
		reverse(arr, 0, v-1)
	}
	return isSorted(arr)
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			return false
		}
	}
	return true
}
