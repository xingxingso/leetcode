package stone_game

import "testing"

func Test_stoneGame(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal0",
			args: args{
				piles: []int{5, 3, 4, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGame(tt.args.piles); got != tt.want {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}

			if got := stoneGameO1(tt.args.piles); got != tt.want {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}

			if got := stoneGameO2(tt.args.piles); got != tt.want {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}

			if got := stoneGameO3(tt.args.piles); got != tt.want {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
