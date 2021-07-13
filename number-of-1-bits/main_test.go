package number_of_1_bits

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				num: 0b00000000000000000000000000001011,
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				num: 0b00000000000000000000000010000000,
			},
			want: 1,
		},
		{
			name: "ex3",
			args: args{
				num: 0b11111111111111111111111111111101,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeightS1(tt.args.num); got != tt.want {
				t.Errorf("hammingWeightS1() = %v, want %v", got, tt.want)
			}

			if got := hammingWeightS2(tt.args.num); got != tt.want {
				t.Errorf("hammingWeightS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
