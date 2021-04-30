package flatten_nested_list_iterator

import (
	"reflect"
	"testing"
)

type args struct {
	nest []*NestedInteger
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
				// [[1,1],2,[1,1]]
				nest: []*NestedInteger{
					&NestedInteger{ns: []*NestedInteger{&NestedInteger{num: 1}, &NestedInteger{num: 1}}},
					&NestedInteger{num: 2},
					&NestedInteger{ns: []*NestedInteger{&NestedInteger{num: 1}, &NestedInteger{num: 1}}},
				},
			},
			want: []int{1, 1, 2, 1, 1},
		},
		{
			name: "ex2",
			args: args{
				// [1,[4,[6]]]
				nest: []*NestedInteger{
					&NestedInteger{num: 1},
					&NestedInteger{
						ns: []*NestedInteger{
							&NestedInteger{num: 4},
							&NestedInteger{ns: []*NestedInteger{&NestedInteger{num: 6}}},
						},
					},
				},
			},
			want: []int{1, 4, 6},
		},
		{
			name: "ex3",
			args: args{
				// [[]]
				nest: []*NestedInteger{
					&NestedInteger{
						ns: []*NestedInteger{},
					},
				},
			},
			want: []int{},
		},
		{
			name: "ex4",
			args: args{
				// [[[8],4]]
				nest: []*NestedInteger{
					&NestedInteger{
						ns: []*NestedInteger{
							&NestedInteger{ns: []*NestedInteger{&NestedInteger{num: 8}}},
							&NestedInteger{num: 4},
						},
					},
				},
			},
			want: []int{8, 4},
		},
	}
}

func TestNestedIterator_Next(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.args.nest)
			got := make([]int, 0)
			for this.HasNext() {
				got = append(got, this.Next())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestedIteratorO1_Next(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			this := ConstructorO1(tt.args.nest)
			got := make([]int, 0)
			for this.HasNext() {
				got = append(got, this.Next())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
