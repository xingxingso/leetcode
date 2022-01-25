package degree_of_an_array

import "testing"

func Test_findShortestSubArray(t *testing.T) {
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
				nums: []int{1, 2, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{1, 2, 2, 3, 1, 4, 2},
			},
			want: 6,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
