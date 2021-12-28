package unique_paths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
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
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "ex2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "ex3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "ex4",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			name: "ex5",
			args: args{
				m: 2,
				n: 2,
			},
			want: 2,
		},
		{
			name: "ex6",
			args: args{ // 超出时限
				m: 51,
				n: 9,
			},
			want: 1916797311,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
			//	t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			//}

			if got := uniquePathsS2(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsS2() = %v, want %v", got, tt.want)
			}

			if got := uniquePathsS3(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsS3() = %v, want %v", got, tt.want)
			}
		})
	}
}
