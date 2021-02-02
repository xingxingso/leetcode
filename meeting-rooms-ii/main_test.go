package meeting_rooms_ii

import "testing"

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			want: 2,
		},
		{
			name: "equal",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			want: 1,
		},
		{
			name: "equal",
			args: args{
				intervals: [][]int{{1, 5}, {8, 9}, {8, 9}},
			},
			want: 2,
		},
		{
			name: "equal",
			args: args{
				intervals: [][]int{{1, 10}, {2, 7}, {3, 19}, {8, 12}, {10, 20}, {11, 30}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}

			if got := minMeetingRoomsO1(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}

			if got := minMeetingRoomsO2(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}

			if got := minMeetingRoomsP1(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
