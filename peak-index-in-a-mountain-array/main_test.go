package peak_index_in_a_mountain_array

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				arr: []int{0, 1, 0},
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				arr: []int{0, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				arr: []int{0, 10, 5, 2},
			},
			want: 1,
		},
		{
			name: "ex4",
			args: args{
				arr: []int{3, 4, 5, 1},
			},
			want: 2,
		},
		{
			name: "ex5",
			args: args{
				arr: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			},
			want: 2,
		},
		{
			name: "ex6",
			args: args{
				arr: []int{3, 5, 3, 2, 0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}

			if got := peakIndexInMountainArrayO1(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArrayO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
