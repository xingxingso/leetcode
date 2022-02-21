package course_schedule_ii

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: []int{0, 1},
		},
		{
			name: "ex2",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0},
					{2, 0},
					{3, 1},
					{3, 2},
				},
			},
			//want: []int{0, 2, 1, 3},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "ex3",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
		{
			name: "ex4",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{0, 1},
					{0, 2},
					{1, 2},
				},
			},
			want: []int{2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
