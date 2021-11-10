package non_decreasing_array

import "testing"

type args struct {
	nums []int
}

func getTests() []struct {
	name string
	args args
	want bool
} {
	return []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				nums: []int{4, 2, 3},
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{4, 2, 1},
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{3, 4, 2, 3},
			},
			want: false,
		},
		{
			name: "ex4",
			args: args{
				nums: []int{5, 7, 1, 8},
			},
			want: true,
		},
	}
}

func Test_checkPossibility(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibilityO1(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibilityO1() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibilityO2(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibilityO2() = %v, want %v", got, tt.want)
			}
		})
	}
}
