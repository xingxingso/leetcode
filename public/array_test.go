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

func Test_intSliceSliceEqual(t *testing.T) {
	type args struct {
		v1 [][]int
		v2 [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				v1: [][]int{{1, 2}, {2, 1}},
				v2: [][]int{{2, 1}, {1, 2}},
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				v1: [][]int{{1, 2, 3}, {1, 3, 2}},
				v2: [][]int{{2, 1, 3}, {2, 3, 1}},
			},
			want: true,
		},
		{
			name: "ex3",
			args: args{
				v1: [][]int{{1, 2, 3}, {4, 5, 6}},
				v2: [][]int{{6, 5, 4}, {3, 2, 1}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intSliceSliceEqual(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("intSliceSliceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intSliceSliceEqual2(t *testing.T) {
	type args struct {
		v1 [][]int
		v2 [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				v1: [][]int{{1, 2}, {2, 1}},
				v2: [][]int{{2, 1}, {1, 2}},
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				v1: [][]int{{1, 2, 3}, {1, 3, 2}},
				v2: [][]int{{2, 1, 3}, {2, 3, 1}},
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				v1: [][]int{{1, 2, 3}, {4, 5, 6}},
				v2: [][]int{{6, 5, 4}, {3, 2, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intSliceSliceEqual2(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("intSliceSliceEqual2() = %v, want %v", got, tt.want)
			}
		})
	}
}
