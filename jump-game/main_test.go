package jump_game

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal0",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "equal1",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
		{
			name: "equal2",
			args: args{
				nums: []int{1, 1, 1, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}

			if got := canJumpS2(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}

			if got := canJumpO1(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}

			if got := canJumpO2(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}

			if got := canJumpO3(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
