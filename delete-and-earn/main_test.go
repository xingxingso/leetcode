package delete_and_earn

import "testing"

func Test_deleteAndEarn(t *testing.T) {
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
				nums: []int{3, 4, 2},
			},
			want: 6,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{2, 2, 3, 3, 3, 4},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteAndEarn(tt.args.nums); got != tt.want {
				t.Errorf("deleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
