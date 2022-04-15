package flood_fill

import (
	"reflect"
	"testing"
)

type args struct {
	image    [][]int
	sr       int
	sc       int
	newColor int
}

type testExamples struct {
	name string
	args args
	want [][]int
}

func getTests() []testExamples {
	return []testExamples{
		{
			name: "ex1",
			args: args{
				image: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "ex1",
			args: args{
				image: [][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
				sr:       0,
				sc:       0,
				newColor: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
		},
	}
}

func Test_floodFill(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
