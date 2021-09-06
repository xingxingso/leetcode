package candy

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				ratings: []int{1, 0, 2}, //2,1,2
			},
			want: 5,
		},
		{
			name: "ex2",
			args: args{
				ratings: []int{1, 2, 2}, // 1,2,1
			},
			want: 4,
		},
		{
			name: "ex3",
			args: args{
				ratings: []int{4, 3, 2, 1}, // 4,3,2,1
			},
			want: 10,
		},
		{
			name: "ex4",
			args: args{
				ratings: []int{1, 3, 2, 2, 1}, // 1,2,1,2,1
			},
			want: 7,
		},
		{
			name: "ex5",
			args: args{
				ratings: []int{1, 2, 3, 1, 0}, // 1,2,3,2,1
			},
			want: 9,
		},
		{
			name: "ex6",
			args: args{
				ratings: []int{29, 51, 87, 87, 72, 12}, // 1,2,3,3,2,1
			},
			want: 12,
		},
		{
			name: "ex7",
			args: args{
				ratings: []int{1, 3, 4, 5, 2}, // 1,2,3,3,2,1
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}

			if got := candyO1(tt.args.ratings); got != tt.want {
				t.Errorf("candyO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
