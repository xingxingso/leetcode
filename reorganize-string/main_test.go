package reorganize_string

import "testing"

func Test_reorganizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				s: "aab",
			},
			want: "aba",
		},
		{
			name: "ex2",
			args: args{
				s: "aaab",
			},
			want: "",
		},
		{
			name: "ex3",
			args: args{
				s: "vvvlo",
			},
			want: "vlvov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeStringO1(tt.args.s); got != tt.want {
				t.Errorf("reorganizeStringO1() = %v, want %v", got, tt.want)
			}

			if got := reorganizeStringO2(tt.args.s); got != tt.want {
				t.Errorf("reorganizeStringO2() = %v, want %v", got, tt.want)
			}
		})
	}
}
