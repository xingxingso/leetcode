package missing_ranges

import (
	"reflect"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "equal",
			args: args{
				nums:  []int{0, 1, 3, 50, 75},
				lower: 0,
				upper: 99,
			},
			want: []string{"2", "4->49", "51->74", "76->99"},
		},
		{
			name: "equal",
			args: args{
				nums:  []int{},
				lower: 0,
				upper: 99,
			},
			want: []string{"0->99"},
		},
		{
			name: "equal",
			args: args{
				nums:  []int{50},
				lower: 0,
				upper: 99,
			},
			want: []string{"0->49", "51->99"},
		},
		{
			name: "equal",
			args: args{
				nums:  []int{},
				lower: 1,
				upper: 1,
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingRanges(tt.args.nums, tt.args.lower, tt.args.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
