package reverse_string

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "ex1",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "ex2",
			args: args{
				s: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverseString(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("%c, reverseString(), want %c", tt.args.s, tt.want)
			}
		})
	}
}
