package russian_doll_envelopes

import "testing"

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal0",
			args: args{
				envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			},
			want: 3,
		},
		{
			name: "equal1",
			args: args{
				envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
