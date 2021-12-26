package ip_to_cidr

import (
	"reflect"
	"testing"
)

func Test_ipToCIDR(t *testing.T) {
	type args struct {
		ip string
		n  int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ex1",
			args: args{
				ip: "255.0.0.7",
				n:  10,
			},
			want: []string{"255.0.0.7/32", "255.0.0.8/29", "255.0.0.16/32"},
		},
		{
			name: "ex2",
			args: args{
				ip: "0.0.0.0",
				n:  1,
			},
			want: []string{"0.0.0.0/32"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ipToCIDR(tt.args.ip, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ipToCIDR() = %v, want %v", got, tt.want)
			}
		})
	}
}
