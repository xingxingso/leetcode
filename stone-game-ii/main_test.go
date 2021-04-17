package stone_game_ii

import "testing"

func Test_stoneGameII(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				piles: []int{2, 7, 9, 4, 4},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIIP1(tt.args.piles); got != tt.want {
				t.Errorf("stoneGameII() = %v, want %v", got, tt.want)
			}
		})
	}
}
