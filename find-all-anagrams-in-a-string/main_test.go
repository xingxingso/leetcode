package find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "equal0",
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			want: []int{0, 6},
		},
		{
			name: "equal1",
			args: args{
				s: "abab",
				p: "ab",
			},
			want: []int{0, 1, 2},
		},
		{
			name: "equal2",
			args: args{
				s: "",
				p: "",
			},
			want: []int{},
		},
		{
			name: "equal3",
			args: args{
				s: "",
				p: "a",
			},
			want: []int{},
		},
		{
			name: "equal4",
			args: args{
				s: "a",
				p: "",
			},
			want: []int{},
		},
		{
			name: "equal5",
			args: args{
				s: "a",
				p: "a",
			},
			want: []int{0},
		},
		{
			name: "equal6",
			args: args{
				s: "zaabcabadaaaefafa",
				p: "aa",
			},
			want: []int{1, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
