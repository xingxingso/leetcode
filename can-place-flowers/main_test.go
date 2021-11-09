package can_place_flowers

import "testing"

type args struct {
	flowerbed []int
	n         int
}

func getTests() []struct {
	name string
	args args
	want bool
} {
	return []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
	}
}

func Test_canPlaceFlowers(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPlaceFlowersO1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowersO1(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowersO1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPlaceFlowersP1(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowersP1(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowersP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
