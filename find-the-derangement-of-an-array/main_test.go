package find_the_derangement_of_an_array

import "testing"

func Test_findDerangement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				n: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDerangementO1(tt.args.n); got != tt.want {
				t.Errorf("findDerangement() = %v, want %v", got, tt.want)
			}

			if got := findDerangementO2(tt.args.n); got != tt.want {
				t.Errorf("findDerangement() = %v, want %v", got, tt.want)
			}

			if got := findDerangementO3(tt.args.n); got != tt.want {
				t.Errorf("findDerangement() = %v, want %v", got, tt.want)
			}
		})
	}
}
