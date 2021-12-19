package beautiful_array

import (
	"testing"
)

func Test_beautifulArray(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				n: 4,
			},
			want: []int{2, 1, 4, 3},
		},
		{
			name: "ex2",
			args: args{
				n: 5,
			},
			want: []int{3, 1, 2, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := beautifulArray(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			if got := beautifulArray(tt.args.n); len(got) != tt.args.n || !isBeautifulArray(got) { // 未验证是 n 的 排列
				t.Errorf("beautifulArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isBeautifulArray(arr []int) bool {
	for i := 1; i < len(arr)-1; i++ {
		for j := i - 1; j >= 0; j-- {
			for k := i + 1; k < len(arr); k++ {
				if arr[i]*2 == arr[j]+arr[k] {
					return false
				}
			}
		}
	}
	return true
}
