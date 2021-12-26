package isomorphic_strings

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				s: "13",
				t: "42",
			},
			want: true,
		},
		{
			name: "ex5",
			args: args{
				s: "ab",
				t: "aa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}

			if got := isIsomorphicP1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphicP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
