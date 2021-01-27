package group_shifted_strings

import (
	"reflect"
	"testing"
)

func Test_groupStrings(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "equal",
			args: args{
				strings: []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
			},
			want: [][]string{{"abc", "bcd", "xyz"}, {"az", "ba"}, {"acef"}, {"a", "z"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupStrings(tt.args.strings); !stringSliceSliceEqual(got, tt.want) {
				t.Errorf("groupStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringSliceSliceEqual(v1, v2 [][]string) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if reflect.DeepEqual(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	return true
}
