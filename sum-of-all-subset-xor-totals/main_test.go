package sum_of_all_subset_xor_totals

import "testing"

func Test_subsetXORSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{1,3},
			},
			want: 6,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{5,1,6},
			},
			want: 28,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{3,4,5,6,7,8},
			},
			want: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetXORSum(tt.args.nums); got != tt.want {
				t.Errorf("subsetXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
