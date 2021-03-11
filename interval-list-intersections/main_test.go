package interval_list_intersections

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		firstList  [][]int
		secondList [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "equal0",
			args: args{
				firstList:  [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
				secondList: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			},
			want: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			name: "equal1",
			args: args{
				firstList:  [][]int{{1, 3}, {5, 9}},
				secondList: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "equal2",
			args: args{
				firstList:  [][]int{},
				secondList: [][]int{{4, 8}, {10, 12}},
			},
			want: [][]int{},
		},
		{
			name: "equal3",
			args: args{
				firstList:  [][]int{{1, 7}},
				secondList: [][]int{{3, 10}},
			},
			want: [][]int{{3, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.firstList, tt.args.secondList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
