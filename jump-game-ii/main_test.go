package jump_game_ii

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}

			if got := jumpO1(tt.args.nums); got != tt.want {
				t.Errorf("jumpO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
