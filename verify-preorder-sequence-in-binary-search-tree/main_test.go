package verify_preorder_sequence_in_binary_search_tree

import "testing"

func Test_verifyPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				preorder: []int{5, 2, 6, 1, 3},
			},
			want: false,
		},
		{
			name: "ex2",
			args: args{
				preorder: []int{5, 2, 1, 3, 6},
			},
			want: true,
		},
		{
			name: "ex3",
			args: args{
				preorder: []int{5, 2, 1, 4, 3},
			},
			want: true,
		},
		{
			name: "ex4",
			args: args{
				preorder: []int{8, 5, 3, 6, 7},
			},
			want: true,
		},
		{
			name: "ex5",
			args: args{
				preorder: []int{50, 25, 13, 18, 38, 30, 75, 60, 55, 70, 90},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPreorder(tt.args.preorder); got != tt.want {
				t.Errorf("verifyPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
