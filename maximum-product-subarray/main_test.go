package maximum_product_subarray

import "testing"

func Test_maxProduct(t *testing.T) {
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
				nums: []int{2, 3, -2, 4},
			},
			want: 6,
		},
		{
			name: "ex2",
			args: args{
				nums: []int{-2, 0, -1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := maxProductS1(tt.args.nums); got != tt.want {
			//	t.Errorf("maxProductS1() = %v, want %v", got, tt.want)
			//}

			if got := maxProductS2(tt.args.nums); got != tt.want {
				t.Errorf("maxProductS2() = %v, want %v", got, tt.want)
			}

			if got := maxProductS3(tt.args.nums); got != tt.want {
				t.Errorf("maxProductS3() = %v, want %v", got, tt.want)
			}

			if got := maxProductO1(tt.args.nums); got != tt.want {
				t.Errorf("maxProductO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
