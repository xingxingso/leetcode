package letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ex1",
			args: args{
				digits: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "ex2",
			args: args{
				digits: "",
			},
			want: []string{},
		},
		{
			name: "ex3",
			args: args{
				digits: "2",
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				//fmt.Printf("fail, %T, %[1]v, %d", tt.want, len(tt.want))
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
