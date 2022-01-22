package reconstruct_itinerary

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	type args struct {
		tickets [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ex1",
			args: args{
				tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			},
			want: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name: "ex2",
			args: args{
				tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			},
			want: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			name: "ex3",
			args: args{
				//map[JFK:[KUL NRT] NRT:[JFK]] 不能先去 KUL
				tickets: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
			},
			want: []string{"JFK", "NRT", "JFK", "KUL"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}

			if got := findItineraryP1(tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItineraryP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
