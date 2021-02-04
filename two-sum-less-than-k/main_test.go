package two_sum_less_than_k

import "testing"

func Test_twoSumLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				nums: []int{34, 23, 1, 24, 75, 33, 54, 8},
				k:    60,
			},
			want: 58,
		},
		{
			name: "equal",
			args: args{
				nums: []int{10, 20, 30},
				k:    15,
			},
			want: -1,
		},
		{
			name: "equal",
			args: args{
				nums: []int{254, 914, 110, 900, 147, 441, 209, 122, 571, 942, 136, 350, 160, 127, 178, 839, 201, 386, 462, 45, 735, 467, 153, 415, 875, 282, 204, 534, 639, 994, 284, 320, 865, 468, 1, 838, 275, 370, 295, 574, 309, 268, 415, 385, 786, 62, 359, 78, 854, 944},
				k:    200,
			},
			want: 198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("twoSumLessThanK() = %v, want %v", got, tt.want)
			}

			if got := twoSumLessThanKP1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("twoSumLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
