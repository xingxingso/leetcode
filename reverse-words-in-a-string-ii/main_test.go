package reverse_words_in_a_string_ii

import (
	"reflect"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "equal",
			args: args{
				s: []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'},
			},
			want: args{
				s: []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'},
			},
		},
		{
			name: "equal",
			args: args{
				s: []byte{},
			},
			want: args{
				s: []byte{},
			},
		},
		{
			name: "equal",
			args: args{
				s: []byte{'t'},
			},
			want: args{
				s: []byte{'t'},
			},
		},
		{
			name: "equal",
			args: args{
				s: []byte{' '},
			},
			want: args{
				s: []byte{' '},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseWords(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want.s) {
				t.Errorf("%c, reverseWords(), want %c", tt.args.s, tt.want.s)
			}
		})
	}
}
