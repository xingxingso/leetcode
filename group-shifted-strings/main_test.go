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
			if got := groupStrings(tt.args.strings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
