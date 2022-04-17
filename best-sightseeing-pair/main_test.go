package best_sightseeing_pair

import "testing"

func Test_maxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				values: []int{8, 1, 5, 2, 6},
			},
			want: 11,
		},
		{
			name: "ex2",
			args: args{
				values: []int{1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreSightseeingPairS1(tt.args.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPairS1() = %v, want %v", got, tt.want)
			}

			if got := maxScoreSightseeingPairO1(tt.args.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPairO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
