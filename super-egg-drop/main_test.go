package super_egg_drop

import "testing"

func Test_superEggDrop(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				k: 1,
				n: 2,
			},
			want: 2,
		},
		{
			name: "equal1",
			args: args{
				k: 2,
				n: 6,
			},
			want: 3,
		},
		{
			name: "equal2",
			args: args{
				k: 3,
				n: 14,
			},
			want: 4,
		},
		{
			name: "equal3",
			args: args{
				k: 4,
				n: 2000,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := superEggDrop(tt.args.k, tt.args.n); got != tt.want {
			// 	t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			// }

			if got := superEggDropO1(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}

			if got := superEggDropO2(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}

			if got := superEggDropO3(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
