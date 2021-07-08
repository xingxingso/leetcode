package public

import (
	"reflect"
	"testing"
)

func Test_copy2(t *testing.T) {
	type args struct {
		dst [][]int
		src [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ex1",
			args: args{
				dst: make([][]int, 2),
				src: [][]int{{1, 2, 3}, {4, 5}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if copy2(tt.args.dst, tt.args.src); !reflect.DeepEqual(tt.args.src, tt.args.dst) {
				t.Errorf("copy2() = %v, want %v", tt.args.dst, tt.args.src)
			}
		})
	}
}
