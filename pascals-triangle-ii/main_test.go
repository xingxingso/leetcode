package pascals_triangle_ii

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				rowIndex: 3,
			},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "ex2",
			args: args{
				rowIndex: 0,
			},
			want: []int{1},
		},
		{
			name: "ex3",
			args: args{
				rowIndex: 1,
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRowO1(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRowO1() = %v, want %v", got, tt.want)
			}

			if got := getRowO2(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRowO2() = %v, want %v", got, tt.want)
			}
		})
	}
}
