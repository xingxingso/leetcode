package burst_balloons

import "testing"

type args struct {
	nums []int
}

type test struct {
	name string
	args args
	want int
}

func getTests() []test {
	tests := []test{
		{
			name: "equal0",
			args: args{
				nums: []int{3, 1, 5, 8},
			},
			want: 167,
		},
		{
			name: "equal1",
			args: args{
				nums: []int{1, 5},
			},
			want: 10,
		},
		{
			name: "equal2",
			args: args{
				nums: []int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3},
			},
			want: 1654,
		},
		{
			name: "equal3",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: 6,
		},
	}
	return tests
}

func Test_maxCoins(t *testing.T) {
	tests := getTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCoinsS2(t *testing.T) {
	tests := getTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoinsS2(tt.args.nums); got != tt.want {
				t.Errorf("maxCoinsS2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCoinsO1(t *testing.T) {
	tests := getTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoinsO1(tt.args.nums); got != tt.want {
				t.Errorf("maxCoinsO1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCoinsO2(t *testing.T) {
	tests := getTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoinsO2(tt.args.nums); got != tt.want {
				t.Errorf("maxCoinsO2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCoinsP1(t *testing.T) {
	tests := getTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoinsP1(tt.args.nums); got != tt.want {
				t.Errorf("maxCoinsP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
