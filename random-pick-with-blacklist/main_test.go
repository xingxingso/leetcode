package random_pick_with_blacklist

import (
	"testing"
)

type args struct {
	n         int
	blacklist []int
}

func getTest() []struct {
	name string
	args args
	want []int
} {
	return []struct {
		name string
		args args
		want []int // 这里仅表示执行次数 和一次执行后的可能正确结果之一
	}{
		{
			name: "ex1",
			args: args{
				n:         1,
				blacklist: []int{},
			},
			want: []int{0, 0, 0},
		},
		{
			name: "ex2",
			args: args{
				n:         2,
				blacklist: []int{},
			},
			want: []int{1, 1, 1},
		},
		{
			name: "ex3",
			args: args{
				n:         3,
				blacklist: []int{1},
			},
			want: []int{0, 0, 2},
		},
		{
			name: "ex4",
			args: args{
				n:         4,
				blacklist: []int{2},
			},
			want: []int{1, 3, 1},
		},
		{
			name: "ex5",
			args: args{
				n:         3,
				blacklist: []int{0, 2},
			},
			want: []int{1, 1, 1},
		},
		{
			name: "ex6",
			args: args{
				n:         2,
				blacklist: []int{},
			},
			want: []int{1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1},
		},
	}
}

func TestSolution(t *testing.T) {
	for _, tt := range getTest() {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor(tt.args.n, tt.args.blacklist)
			for _, v := range tt.want {
				// 随机性没有验证
				got := s.Pick()
				if !(got == v || (got >= 0 && got < tt.args.n && !inArray(tt.args.blacklist, got))) {
					t.Errorf("Pick() = %v, want %v", got, v)
				}
			}
		})
	}
}

func inArray(nums []int, n int) bool {
	for _, v := range nums {
		if v == n {
			return true
		}
	}
	return false
}
