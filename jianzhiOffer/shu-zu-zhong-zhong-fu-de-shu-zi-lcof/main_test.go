package shu_zu_zhong_zhong_fu_de_shu_zi_lcof

import "testing"

type args struct {
	nums []int
}

func getTests() []struct {
	name string
	args args
	want map[int]bool
} {
	return []struct {
		name string
		args args
		want map[int]bool
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: map[int]bool{2: true, 3: true},
		},
	}
}

func Test_findRepeatNumber(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); !tt.want[got] {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
