package stone_game_v

import "testing"

func Test_stoneGameV(t *testing.T) {
	type args struct {
		stoneValue []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				stoneValue: []int{6, 2, 3, 4, 5, 5},
			},
			want: 18,
		},
		{
			name: "equal1",
			args: args{
				stoneValue: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 28,
		},
		{
			name: "equal2",
			args: args{
				stoneValue: []int{4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameV(tt.args.stoneValue); got != tt.want {
				t.Errorf("stoneGameV() = %v, want %v", got, tt.want)
			}
		})
	}
}
