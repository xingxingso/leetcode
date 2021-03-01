package best_time_to_buy_and_sell_stock_iv

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
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
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "equal1",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
		{
			name: "equal2",
			args: args{
				k:      1,
				prices: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "equal3",
			args: args{
				k:      2,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			},
			want: 13,
		},
		{
			name: "equal4",
			args: args{
				k:      2,
				prices: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}

			if got := maxProfitO1(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}

			if got := maxProfitO2(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
