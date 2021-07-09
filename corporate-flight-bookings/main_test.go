package corporate_flight_bookings

import (
	"reflect"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {
	type args struct {
		bookings [][]int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				bookings: [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
				n:        5,
			},
			want: []int{10, 55, 45, 25, 25},
		},
		{
			name: "ex2",
			args: args{
				bookings: [][]int{{1, 2, 10}, {2, 2, 15}},
				n:        2,
			},
			want: []int{10, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookingsS1(tt.args.bookings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookingsS1() = %v, want %v", got, tt.want)
			}

			if got := corpFlightBookingsP1(tt.args.bookings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookingsP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
