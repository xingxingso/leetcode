package stone_game_iii

import "testing"

func Test_stoneGameIII(t *testing.T) {
	type args struct {
		stoneValue []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "equal0",
			args: args{
				stoneValue: []int{1, 2, 3, 7},
			},
			want: "Bob",
		},
		{
			name: "equal1",
			args: args{
				stoneValue: []int{1, 2, 3, -9},
			},
			want: "Alice",
		},
		{
			name: "equal2",
			args: args{
				stoneValue: []int{1, 2, 3, 6},
			},
			want: "Tie",
		},
		{
			name: "equal3",
			args: args{
				stoneValue: []int{1, 2, 3, -1, -2, -3, 7},
			},
			want: "Alice",
		},
		{
			name: "equal4",
			args: args{
				stoneValue: []int{-1, -2, -3},
			},
			want: "Tie",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIII(tt.args.stoneValue); got != tt.want {
				t.Errorf("stoneGameIII() = %v, want %v", got, tt.want)
			}

			if got := stoneGameIIIO1(tt.args.stoneValue); got != tt.want {
				t.Errorf("stoneGameIII() = %v, want %v", got, tt.want)
			}

			if got := stoneGameIIIO2(tt.args.stoneValue); got != tt.want {
				t.Errorf("stoneGameIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
