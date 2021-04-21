package stone_game_vi

import "testing"

func Test_stoneGameVI(t *testing.T) {
	type args struct {
		aliceValues []int
		bobValues   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				aliceValues: []int{1, 3},
				bobValues:   []int{2, 1},
			},
			want: 1,
		},
		{
			name: "equal1",
			args: args{
				aliceValues: []int{1, 2},
				bobValues:   []int{3, 1},
			},
			want: 0,
		},
		{
			name: "equal2",
			args: args{
				aliceValues: []int{2, 4, 3},
				bobValues:   []int{1, 6, 7},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameVI(tt.args.aliceValues, tt.args.bobValues); got != tt.want {
				t.Errorf("stoneGameVI() = %v, want %v", got, tt.want)
			}
		})
	}
}
