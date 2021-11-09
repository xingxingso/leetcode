package queue_reconstruction_by_height

import (
	"reflect"
	"testing"
)

type args struct {
	people [][]int
}

func getTests() []struct {
	name string
	args args
	want [][]int
} {
	return []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ex1",
			args: args{
				people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			},
			want: [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
	}
}

func Test_reconstructQueue(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}

	//for _, tt := range getTests() {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := reconstructQueueS1(tt.args.people); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("reconstructQueueS1() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
