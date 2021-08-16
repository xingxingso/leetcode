package rotate_array

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
	k    int
}

func getTests() []struct {
	name string
	args args
	want []int
} {
	return []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "ex2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "ex3",
			args: args{
				nums: func() []int {
					nums := make([]int, 20000000)
					for i := 0; i < 20000000; i++ {
						nums[i] = i
					}
					return nums
				}(),
				k: 16000000,
			},
			want: func() []int {
				nums := make([]int, 20000000)
				for i := 0; i < 20000000; i++ {
					nums[i] = (i + 4000000) % 20000000
				}
				return nums
			}(),
		},
	}
}

func Test_rotateP1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if rotateP1(tt.args.nums, tt.args.k); !reflect.DeepEqual(tt.args.nums, tt.want) {
				if tt.name != "ex3" {
					t.Errorf("rotateP1() got %v, want %v", tt.args.nums, tt.want)
				} else {
					t.Error("rotateP1() error")
				}
			}
		})
	}
}

func Test_rotateP2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if rotateP2(tt.args.nums, tt.args.k); !reflect.DeepEqual(tt.args.nums, tt.want) {
				if tt.name != "ex3" {
					t.Errorf("rotateP2() got %v, want %v", tt.args.nums, tt.want)
				} else {
					t.Error("rotateP2() error")
				}
			}
		})
	}
}
