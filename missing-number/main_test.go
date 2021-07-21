package missing_number

import "testing"

func Test_missingNumber(t *testing.T) {
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
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
		{
			name: "ex4",
			args: args{
				nums: []int{0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}

			if got := missingNumberP1(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberP1() = %v, want %v", got, tt.want)
			}

			if got := missingNumberP2(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberP2() = %v, want %v", got, tt.want)
			}
		})
	}
}
