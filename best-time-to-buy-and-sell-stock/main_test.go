package best_time_to_buy_and_sell_stock

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "equal1",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "equal2",
			args: args{
				prices: []int{414, 863, 393, 674, 205, 793, 229, 379, 37, 455, 594, 36, 312, 667, 441, 411, 514, 344, 681,
					359, 865, 124, 984, 670, 509, 337, 495, 266, 275, 356, 26, 229, 51, 557, 292, 975, 551, 985, 445, 710,
					467, 31, 890, 694, 127, 349, 631, 322, 595, 59, 433, 173, 944, 305, 662, 379, 864, 835, 355, 411, 506,
					10, 716, 918, 872, 716, 887, 453, 706, 416, 245, 611, 6, 403, 894, 94, 852, 733, 890, 131, 481, 723, 571,
					335, 695, 475},
			},
			want: 959,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}

			if got := maxProfitS1(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}

			if got := maxProfitS0(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}

			if got := maxProfitO1(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
