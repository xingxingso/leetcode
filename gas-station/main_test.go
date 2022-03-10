package gas_station

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			want: -1,
		},
		{
			name: "ex3",
			args: args{
				gas:  []int{5, 1, 2, 3, 4},
				cost: []int{4, 4, 1, 5, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuitS1(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuitS1() = %v, want %v", got, tt.want)
			}

			if got := canCompleteCircuitS2(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuitS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
