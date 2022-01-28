package advantage_shuffle

import (
	"reflect"
	"testing"
)

func Test_advantageCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				nums1: []int{2, 7, 11, 15},
				nums2: []int{1, 10, 4, 11},
			},
			want: []int{2, 11, 7, 15},
		},
		{
			name: "ex2",
			args: args{
				nums1: []int{12, 24, 8, 32},
				nums2: []int{13, 25, 32, 11},
			},
			want: []int{24, 32, 8, 12},
		},
		{
			name: "ex3",
			args: args{
				nums1: []int{9, 1, 2, 4, 5},
				nums2: []int{6, 2, 9, 1, 4},
			},
			want: []int{9, 4, 1, 2, 5},
		},
		{
			name: "ex4",
			args: args{
				nums1: []int{9, 1, 5, 4, 2, 5},
				nums2: []int{6, 4, 9, 1, 4, 2},
			},
			want: []int{9, 5, 1, 2, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := advantageCount(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("advantageCount() = %v, want %v", got, tt.want)
			}

			if got := advantageCountS2(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("advantageCountS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
