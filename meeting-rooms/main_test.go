package meeting_rooms

import "testing"

func Test_canAttendMeetings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				intervals: [][]int{{8, 11}, {17, 20}, {17, 20}},
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				intervals: [][]int{{13, 15}, {1, 13}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetings(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}
