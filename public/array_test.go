package public

import (
	"reflect"
	"testing"
)

func Test_copy2(t *testing.T) {
	type args struct {
		dst [][]int
		src [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ex1",
			args: args{
				dst: make([][]int, 2),
				src: [][]int{{1, 2, 3}, {4, 5}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if copy2(tt.args.dst, tt.args.src); !reflect.DeepEqual(tt.args.src, tt.args.dst) {
				t.Errorf("copy2() = %v, want %v", tt.args.dst, tt.args.src)
			}
		})
	}
}

func Test_sliceSameValue(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				nums1: []int{2, 2, 3, 1},
				nums2: []int{1, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "ex2",
			args: args{
				nums1: []int{2, 1, 2, 3, 1},
				nums2: []int{1, 3, 2, 1, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceSameValue(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("sliceSameValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
