package capacity_to_ship_packages_within_d_days

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		D       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				D:       5,
			},
			want: 15,
		},
		{
			name: "ex2",
			args: args{
				weights: []int{3, 2, 2, 4, 1, 4},
				D:       3,
			},
			want: 6,
		},
		{
			name: "ex3",
			args: args{
				weights: []int{1, 2, 3, 1, 1},
				D:       4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.D); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}

			if got := shipWithinDaysO1(tt.args.weights, tt.args.D); got != tt.want {
				t.Errorf("shipWithinDaysO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
